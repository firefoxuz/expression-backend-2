package middleware

import (
	"context"
	"expression-backend/internal/entities"
	"expression-backend/internal/utils"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorizationToken := r.Header.Get("authorization")

		if authorizationToken == "" {
			_ = utils.RespondWith401(w, "jwt token is not set on Authorization header")
			return
		}

		if len(strings.Split(r.Header.Get("Authorization"), " ")) != 2 {
			_ = utils.RespondWith401(w, "invalid authorization token")
			return
		}

		tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]

		if tokenString == "" {
			_ = utils.RespondWith401(w, "unauthorized")
			return
		}

		login, err := utils.ValidateToken(tokenString)
		if err != nil {
			_ = utils.RespondWith401(w, "unauthorized")
			return
		}

		user, err := entities.FindUserByLogin(login)

		if err != nil {
			_ = utils.RespondWith500(w)
			return
		}

		ctx := r.Context()

		ctx = context.WithValue(ctx, "login", login)
		ctx = context.WithValue(ctx, "user", user)
		req := r.WithContext(ctx)

		next.ServeHTTP(w, req)
	}
}
