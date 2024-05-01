package middleware

import (
	"github.com/AliSahib998/QuotesAssesments/model"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"strings"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authIgnore := []string{"/user/register", "/user/login", "/health"}
		requestPath := r.URL.Path

		for _, value := range authIgnore {

			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")

		if tokenHeader == "" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(403)
			return
		}

		tokenSplit := strings.Split(tokenHeader, " ")
		if len(tokenSplit) != 2 {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(403)
			return
		}

		if tokenSplit[0] != "Bearer" {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(403)
			return
		}

		jwtToken := tokenSplit[1]
		tk := &model.Claim{}

		token, err := jwt.ParseWithClaims(jwtToken, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN_KEY")), nil
		})

		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(403)
			return
		}

		if !token.Valid {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(403)
			return
		}

		next.ServeHTTP(w, r)
	})
}
