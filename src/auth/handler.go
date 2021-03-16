package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/auth/database"
	"github.com/loopcontext/auth-api-go/src/auth/models"
	"github.com/loopcontext/auth-api-go/src/auth/rest"
	"github.com/loopcontext/auth-api-go/src/utils"
	"github.com/markbates/goth/gothic"
	"github.com/rs/zerolog/log"
)

var (
	errEmailPass = errors.New("please enter a correct email address and password")
	domain       = "app"
)

// Register entry point of the slsfn /v{X}/auth/register
func Register(db *gen.DB) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var input models.RegisterInput
		err := rest.ReadJSON(res, req, &input)
		if err != nil {
			rest.HandleErr(res, req, err)

			return
		}
		err = input.Validate()
		if err != nil {
			rest.HandleErr(res, req, err)

			return
		}
		log.Debug().Msgf("Registering user: %+v", input)
		passwd, err := utils.HashPassword(input.Password)
		if err != nil {
			rest.HandleErr(res, req, err)

			return
		}
		user := &gen.User{
			Active:      true,
			DisplayName: utils.StrToPtrStr(input.FirstName + " " + input.LastName),
			FirstName:   utils.StrToPtrStr(input.FirstName),
			LastName:    utils.StrToPtrStr(input.LastName),
			Password:    &passwd,
			Email:       input.Email,
		}
		resp := rest.Response{Status: http.StatusUnprocessableEntity}
		if db.Query().Model(user).First(user, "email = ?", input.Email).Error != nil {
			user.ID = uuid.Must(uuid.NewV4()).String()
			if err := db.Query().Model(user).Create(user).Error; err != nil {
				status := []int{}
				if errors.Is(err, gorm.ErrRecordNotFound) {
					status = append(status, http.StatusNotFound)
				}
				rest.HandleErr(res, req, err, status...)

				return
			}
			userProfile := &gen.Profile{
				ID:             uuid.Must(uuid.NewV4()).String(),
				Email:          user.Email,
				ExternalUserID: &user.ID,
				Provider:       &domain,
				Name:           user.DisplayName,
				FirstName:      user.FirstName,
				LastName:       user.LastName,
				Users:          []*gen.User{user},
			}
			if err := db.Query().Model(userProfile).Create(userProfile).Error; err != nil {
				status := []int{}
				if errors.Is(err, gorm.ErrRecordNotFound) {
					status = append(status, http.StatusNotFound)
				}
				rest.HandleErr(res, req, err, status...)

				return
			}

			resp.Msg = "successfully created"
			resp.Status = http.StatusCreated
		} else {
			// already exists
			resp.Msg = "already exists"
		}
		err = rest.SendJSON(res, req, resp, http.StatusCreated)
		if err != nil {
			rest.HandleErr(res, req, err)

			return
		}
	}
}

// Login entry point of the slsfn /v{X}/auth/login
func Login(db *gen.DB) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		var input models.LoginInput
		err := rest.ReadJSON(res, req, &input)
		if err != nil {
			rest.HandleErr(res, req, err)

			return
		}
		err = input.Validate()
		if err != nil {
			rest.HandleErr(res, req, err)

			return
		}
		log.Debug().Msgf("Login user: %s", input.Email)
		user := &gen.User{}
		if err := db.Query().Model(user).Preload("Roles").Preload("Roles.Permissions").Preload("Permissions").
			First(user, "email = ? AND active = ?", input.Email, true).Error; err != nil {
			status := []int{}
			if errors.Is(err, gorm.ErrRecordNotFound) {
				status = append(status, http.StatusNotFound)
			}
			rest.HandleErr(res, req, errEmailPass, status...)

			return
		}
		passw := ""
		if user.Password != nil {
			passw = *user.Password
		}
		if err := utils.PasswordCheck(passw, input.Password); err != nil {
			rest.HandleErr(res, req, errEmailPass, http.StatusNotFound)

			return
		}
		now := time.Now()
		token, err := generateToken(user, jwt.StandardClaims{
			Id: user.ID, Issuer: "app", Subject: user.ID, Audience: req.Host,
			ExpiresAt: now.Add(12 * time.Hour).UTC().Unix(), IssuedAt: now.UTC().Unix(),
			NotBefore: now.UTC().Unix(),
		})
		if err != nil {
			rest.HandleErr(res, req, err)

			return
		}
		response := map[string]interface{}{
			"type":  "Bearer",
			"token": token,
		}
		if err := rest.SendJSON(res, req, response); err != nil {
			rest.HandleErr(res, req, err)
		}
	}
}

// Begin entry point of the slsfn /v{X}/auth/[provider]
func Begin(res http.ResponseWriter, req *http.Request) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	req = addProviderToContext(req, mux.Vars(req)[string(utils.ProjectContextKeys.ProviderCtxKey)])
	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(res, req); err != nil {
		log.Error().Msg(err.Error())
		gothic.BeginAuthHandler(res, req)
	}
}

// BeginNamedProvider entry point of the slsfn /v{X}/auth/[provider] - Named helper
func BeginNamedProvider(provider string, res http.ResponseWriter, req *http.Request) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	req = addProviderToContext(req, provider)
	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(res, req); err != nil {
		log.Error().Msg(err.Error())
		gothic.BeginAuthHandler(res, req)
	}
}

// CallbackHandler to complete auth provider flow
func CallbackHandler(db *gen.DB) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// You have to add value context with provider name to get provider name in GetProviderName method
		req = addProviderToContext(req, mux.Vars(req)[string(utils.ProjectContextKeys.ProviderCtxKey)])
		callbackHandlerExec(res, req, db)
	}
}

// CallbackHandlerNamedProvider to complete auth provider flow - Named provider helper
func CallbackHandlerNamedProvider(db *gen.DB, provider string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		req = addProviderToContext(req, provider)
		callbackHandlerExec(res, req, db)
	}
}

func callbackHandlerExec(res http.ResponseWriter, req *http.Request, db *gen.DB) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		rest.HandleErr(res, req, err, http.StatusUnprocessableEntity)

		return
	}
	u, err := database.FindUserByJWT(db, user.UserID, user.Email, user.Provider)
	if err != nil {
		if u, err = database.UpsertUserProfile(db, &user); err != nil {
			log.Err(fmt.Errorf("[Auth.CallbackHandler.Error]: %q", err)).Send()
			rest.HandleErr(res, req, err, http.StatusUnprocessableEntity)
		}
	}
	token, err := generateToken(u, jwt.StandardClaims{
		Id: user.UserID, Issuer: user.Provider, Subject: u.ID, Audience: req.Host,
		ExpiresAt: user.ExpiresAt.UTC().Unix(), IssuedAt: time.Now().UTC().Unix(),
		NotBefore: time.Now().UTC().Unix(),
	})
	if err != nil {
		rest.HandleErr(res, req, err, http.StatusUnprocessableEntity)

		return
	}
	response := map[string]interface{}{
		"type":  "Bearer",
		"token": token,
	}
	if err := rest.SendJSON(res, req, response); err != nil {
		rest.HandleErr(res, req, err)
	}
}

// Logout logs out of the auth provider
func Logout(res http.ResponseWriter, req *http.Request) {
	req = addProviderToContext(req, mux.Vars(req)[string(utils.ProjectContextKeys.ProviderCtxKey)])
	err := gothic.Logout(res, req)
	if err != nil {
		rest.HandleErr(res, req, err)

		return
	}
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

// Get the user roles from object
func getUserRoles(u *gen.User) (roles []string) {
	for _, r := range u.Roles {
		roles = append(roles, r.Name)
	}

	return
}

func generateToken(user *gen.User, stdClaims jwt.StandardClaims) (string, error) {
	claims := gen.JWTClaims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			Id:        stdClaims.Id,
			Subject:   stdClaims.Subject,
			Issuer:    stdClaims.Issuer,
			Audience:  stdClaims.Audience,
			IssuedAt:  time.Now().UTC().Unix(),
			NotBefore: time.Now().UTC().Unix(),
			ExpiresAt: stdClaims.ExpiresAt,
		}, Roles: getUserRoles(user),
		Permissions: getUserPermissions(user),
	}
	if user.DisplayName != nil {
		claims.Name = *user.DisplayName
	}
	if user.NickName != nil {
		claims.Nickname = *user.NickName
	}
	if user.AvatarURL != nil {
		claims.Picture = *user.AvatarURL
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(utils.MustGet("AUTH_JWT_SIGNING_ALGORITHM")), claims)

	return jwtToken.SignedString([]byte(utils.MustGet("AUTH_JWT_SECRET")))
}

// Get user's permissions from object
func getUserPermissions(u *gen.User) map[string]string {
	perms := make(map[string]string)
	for _, p := range u.Permissions {
		charIdx := strings.Index(p.Tag, ":")
		if perms[p.Tag[:charIdx]] == "" {
			perms[p.Tag[:charIdx]] += p.Tag[charIdx+1 : charIdx+2]
		} else {
			perms[p.Tag[:charIdx]] += "," + p.Tag[charIdx+1:charIdx+2]
		}
	}

	return perms
}
