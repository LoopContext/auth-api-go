package auth

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"

	"github.com/loopcontext/auth-api-go/gen"
	"github.com/loopcontext/auth-api-go/src/auth/database"
	"github.com/loopcontext/auth-api-go/src/auth/utils"

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

// CallbackHandler to complete auth provider flow
func CallbackHandler(db *gen.DB) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		// You have to add value context with provider name to get provider name in GetProviderName method
		req = addProviderToContext(req, mux.Vars(req)[string(utils.ProjectContextKeys.ProviderCtxKey)])
		callbackHandlerExec(res, req, db)
	}
}

// CallbackHandlerVercel to complete auth provider flow
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
			Email:  user.Email,
			Avatar: user.AvatarURL,
			StandardClaims: jwt.StandardClaims{
				Id:        user.UserID,
				Issuer:    user.Provider,
				Subject:   user.Email,
				Audience:  req.Host,
				IssuedAt:  time.Now().UTC().Unix(),
				NotBefore: time.Now().UTC().Unix(),
				ExpiresAt: user.ExpiresAt.UTC().Unix(),
			},
			Roles: getJWTRoles(u),
		})
	token, err := jwtToken.SignedString([]byte(utils.MustGet("AUTH_JWT_SECRET")))
	if err != nil {
		abortWithError(&res, http.StatusInternalServerError, err)
		return
	}
	jwtToken = jwt.NewWithClaims(jwt.GetSigningMethod(utils.MustGet("AUTH_JWT_SIGNING_ALGORITHM")),
		gen.JWTClaims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				Id:        user.UserID,
				Issuer:    user.Provider,
				Subject:   user.Email,
				Audience:  req.Host,
				IssuedAt:  time.Now().UTC().Unix(),
				NotBefore: time.Now().UTC().Unix(),
				ExpiresAt: user.ExpiresAt.Add(2 * time.Hour).UTC().Unix(),
			},
		})
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

func getJWTRoles(u *gen.User) (roles []gen.JWTRole) {
	var d string
	for _, r := range u.Roles {
		if r.Description != nil {
			d = *r.Description
		}
		roles = append(roles, gen.JWTRole{
			Name:        r.Name,
			Description: d,
			Permissions: getJWTRolePermissions(r),
		})
	}
	return
}

func getJWTRolePermissions(r *gen.Role) (permissions []string) {
	for _, p := range r.Permissions {
		permissions = append(permissions, p.Tag)
	}
	return
}
