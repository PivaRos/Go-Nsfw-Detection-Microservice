package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pivaros/go-image-recognition/constants"
	"github.com/pivaros/go-image-recognition/utils"
)

func Authenticate(roles []constants.Role, app *utils.AppState) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			token, err := jwt.ParseWithClaims(tokenString, &utils.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
				return app.Env.Jwt_Secret_Key, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if claim, ok := token.Claims.(*utils.UserClaims); ok && token.Valid {
				var found bool = false
				for _, value := range roles {
					if value == claim.Role {
						found = true
					}
				}
				var id, gov_id, role string
				if found {
					//make query to find the user
					err = app.Db.QueryRow("SELECT id, gov_id, role FROM users WHERE access_token = $1", tokenString).Scan(&id, &gov_id, &role)
					if err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						return
					}
					//assign the needed values to the context
					ctx := context.WithValue(r.Context(), constants.UserIdContextKey, id)
					r = r.WithContext(ctx)
					ctx = context.WithValue(r.Context(), constants.UserGovIdContextKey, gov_id)
					r = r.WithContext(ctx)
					ctx = context.WithValue(r.Context(), constants.UserRoleContextKey, role)
					r = r.WithContext(ctx)
					//move to the next handler
					next.ServeHTTP(w, r)
					return
				}
				w.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		})
	}
}
