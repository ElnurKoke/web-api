package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ElnurKoke/web-api.git/internal/model"
)

func (h *Server) login(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodPost:
		var User model.User
		err := json.NewDecoder(r.Body).Decode(&User)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, expired, err := h.store.CheckUser(model.User{
			Username: User.Username,
			Password: User.Password,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Path:    "/",
			Expires: expired,
		})

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "You have successfully login, your token:%s ", token)

	case http.MethodGet:
		fmt.Fprintf(w, "Welcome to login page")
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *Server) logOut(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logout" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := h.store.DeleteToken(c.Value); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name: "token",

		MaxAge: -1,
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You have successfully logged out of your account, the token has been deleted")
}
