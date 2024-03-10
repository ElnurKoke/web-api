package apiserver

import (
	"context"
	"net/http"
	"time"

	"github.com/ElnurKoke/web-api.git/internal/model"
)

func (h *Server) middleWareGetUser(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		c, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", model.User{})))
			return
		}
		user, err = h.store.GetUserByToken(c.Value)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", model.User{})))
			return
		}
		if user.ExpiresAt.Before(time.Now()) {
			if err := h.store.DeleteToken(c.Value); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			http.Error(w, "Token expired", http.StatusOK)
			return
		}
		user.IsAuth = true

		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", user)))

	}
}
