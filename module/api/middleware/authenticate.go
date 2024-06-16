package middleware

import (
	"context"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/pivaros/go-image-recognition/constants"
	"github.com/pivaros/go-image-recognition/structs"
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
				if found {
					var user structs.User
					err = app.Db.QueryRow("SELECT id, first_name, last_name, phone, gov_id, role, access_token FROM users WHERE access_token = $1", tokenString).Scan(
						&user.Id, &user.FirstName, &user.LastName, &user.Phone, &user.GovId, &user.Role, &user.AccessToken)
					if err != nil {
						w.WriteHeader(http.StatusUnauthorized)
						return
					}

					ctx := context.WithValue(r.Context(), constants.UserDataContextKey, user)
					r = r.WithContext(ctx)
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
