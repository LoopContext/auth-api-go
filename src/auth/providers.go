package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/loopcontext/auth-api-go/src/utils"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/auth0"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/twitter"
)

func addProviderToContext(r *http.Request, provider interface{}) *http.Request {
	// gothic.Store = sessions.NewCookieStore([]byte("<your secret here>"))
	checkProviders(provider.(string))
	return r.WithContext(context.WithValue(r.Context(),
		string(utils.ProjectContextKeys.GothicProviderCtxKey), provider))
}

func parseScopes(scopes string) []string {
	return strings.Split(scopes, ",")
}

func checkProviders(provider string) {
	provisioned := false
	for _, p := range goth.GetProviders() {
		if provisioned = provider == p.Name(); provisioned {
			break
		}
	}
	if !provisioned {
		switch provider {
		case "google":
			goth.UseProviders(google.New(utils.MustGet("PROVIDER_GOOGLE_KEY"), utils.MustGet("PROVIDER_GOOGLE_SECRET"),
				fmt.Sprintf(utils.MustGet("AUTH_JWT_CALLBACK"), provider), parseScopes(utils.MustGet("PROVIDER_GOOGLE_SCOPES"))...))
		case "auth0":
			goth.UseProviders(auth0.New(utils.MustGet("PROVIDER_AUTH0_KEY"), utils.MustGet("PROVIDER_AUTH0_SECRET"),
				fmt.Sprintf(utils.MustGet("AUTH_JWT_CALLBACK"), provider), utils.MustGet("PROVIDER_AUTH0_DOMAIN"),
				parseScopes(utils.MustGet("PROVIDER_AUTH0_SCOPES"))...))
		case "facebook":
			goth.UseProviders(facebook.New(utils.MustGet("PROVIDER_FACEBOOK_KEY"), utils.MustGet("PROVIDER_FACEBOOK_SECRET"),
				fmt.Sprintf(utils.MustGet("AUTH_JWT_CALLBACK"), provider), parseScopes(utils.MustGet("PROVIDER_FACEBOOK_SCOPES"))...))
		case "twitter":
			goth.UseProviders(twitter.New(utils.MustGet("PROVIDER_TWITTER_KEY"), utils.MustGet("PROVIDER_TWITTER_SECRET"),
				fmt.Sprintf(utils.MustGet("AUTH_JWT_CALLBACK"), provider)))
		}
	}
}
