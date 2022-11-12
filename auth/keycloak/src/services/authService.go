package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak"
)

type LoginResponse struct {
	AccessToken string `json:"access_token"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var (
	clientId     = "go-auth-test"
	clientSecret = "somesecret"
	realm        = "go-auth"
	hostname     = "http://localhost:8180"
)

var client gocloak.GoCloak

func InitializeOauthServer() {
	client = gocloak.NewClient(hostname)
}

func Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if len(authHeader) < 1 {
			w.WriteHeader(401)
			fmt.Fprint(w, "authorization header not found")
			return
		}

		accessTokenList := strings.Split(authHeader, " ")
		if len(accessTokenList) <= 1 {
			fmt.Fprint(w, "no auth header")
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]
		rptToken, err := client.RetrospectToken(accessToken, clientId, clientSecret, realm)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprint(w, err.Error())
			return
		}

		isTokenValid := rptToken.Active
		if !isTokenValid {
			w.WriteHeader(401)
			fmt.Fprint(w, "token is invalid")
			return
		}

		next.ServeHTTP(w, r)
	})
}
