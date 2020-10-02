package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/auth/database"
	"github.com/loopcontext/auth-api-go/src/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/markbates/goth/gothic"
)

// Begin entry point of the slsfn /v{X}/auth/[provider]
func Begin(res http.ResponseWriter, req *http.Request) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	req = addProviderToContext(req, mux.Vars(req)[string(utils.ProjectContextKeys.ProviderCtxKey)])
	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(res, req); err != nil {
		gothic.BeginAuthHandler(res, req)
	}
}

// BeginVercel entry point of the slsfn /v{X}/auth/[provider] - Vercel helper
func BeginVercel(provider string, res http.ResponseWriter, req *http.Request) {
	// You have to add value context with provider name to get provider name in GetProviderName method
	req = addProviderToContext(req, provider)
	// try to get the user without re-authenticating
	if _, err := gothic.CompleteUserAuth(res, req); err != nil {
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

// CallbackHandlerVercel to complete auth provider flow - Vercel helper
func CallbackHandlerVercel(db *gen.DB, provider string) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		req = addProviderToContext(req, provider)
		callbackHandlerExec(res, req, db)
	}
}

func callbackHandlerExec(res http.ResponseWriter, req *http.Request, db *gen.DB) {
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		abortWithError(&res, http.StatusInternalServerError, err)
		return
	}
	u, err := database.FindUserByJWT(db, user.UserID, user.Email, user.Provider)
	if err != nil {
		if u, err = database.UpsertUserProfile(db, &user); err != nil {
			log.Err(fmt.Errorf("[Auth.CallbackHandler.Error]: %q", err)).Send()
			abortWithError(&res, http.StatusInternalServerError, err)
		}
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod(utils.MustGet("AUTH_JWT_SIGNING_ALGORITHM")),
		gen.JWTClaims{
			Name:     user.Name,
			Nickname: user.NickName,
			Email:    user.Email,
			Picture:  user.AvatarURL,
			StandardClaims: jwt.StandardClaims{
				Id:        user.UserID,
				Subject:   u.ID,
				Issuer:    user.Provider,
				Audience:  req.Host,
				IssuedAt:  time.Now().UTC().Unix(),
				NotBefore: time.Now().UTC().Unix(),
				ExpiresAt: user.ExpiresAt.UTC().Unix(),
			},
			Roles:       getUserRoles(u),
			Permissions: getUserPermissions(u),
		})
	token, err := jwtToken.SignedString([]byte(utils.MustGet("AUTH_JWT_SECRET")))
	if err != nil {
		abortWithError(&res, http.StatusInternalServerError, err)
		return
	}
	// jwtToken = jwt.NewWithClaims(jwt.GetSigningMethod(utils.MustGet("AUTH_JWT_SIGNING_ALGORITHM")),
	// 	gen.JWTClaims{
	// 		Email:   user.Email,
	// 		Picture: user.AvatarURL,
	// 		StandardClaims: jwt.StandardClaims{
	// 			Id:        user.UserID,
	// 			Issuer:    user.Provider,
	// 			Subject:   user.Email,
	// 			Audience:  req.Host,
	// 			IssuedAt:  time.Now().UTC().Unix(),
	// 			NotBefore: time.Now().UTC().Unix(),
	// 			ExpiresAt: user.ExpiresAt.Add(2 * time.Hour).UTC().Unix(),
	// 		},
	// 	})
	// refreshtoken, err := jwtToken.SignedString([]byte(utils.MustGet("AUTH_JWT_SECRET")))
	response := map[string]interface{}{
		"type":  "Bearer",
		"token": token,
		// "refresh_token": refreshtoken,
	}
	parseJSON(&res, http.StatusOK, response)
}

// Logout logs out of the auth provider
func Logout(res http.ResponseWriter, req *http.Request) {
	req = addProviderToContext(req, mux.Vars(req)[string(utils.ProjectContextKeys.ProviderCtxKey)])
	gothic.Logout(res, req)
	res.Header().Set("Location", "/")
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func abortWithError(w *http.ResponseWriter, sc int, err error) {
	(*w).WriteHeader(sc)
}

// parseJSON friday 13th
func parseJSON(w *http.ResponseWriter, sc int, j interface{}) (err error) {
	b, err := json.Marshal(j)
	fmt.Fprint(*w, string(b))
	return err
}

func getUserRoles(u *gen.User) (roles []string) {
	for _, r := range u.Roles {
		roles = append(roles, r.Name)
	}
	return
}

func getUserPermissions(u *gen.User) map[string]string {
	perms := make(map[string]string)
	for _, p := range u.Permissions {
		c := strings.Index(p.Tag, ":")
		if perms[p.Tag[:c]] == "" {
			perms[p.Tag[:c]] += p.Tag[c+1 : c+2]
		} else {
			perms[p.Tag[:c]] += "," + p.Tag[c+1:c+2]
		}
	}
	return perms
}
