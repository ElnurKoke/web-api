package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ElnurKoke/web-api.git/internal/model"
	"github.com/ElnurKoke/web-api.git/internal/store"
)

func (h *Server) register(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodPost:
		var newUser model.User
		err := json.NewDecoder(r.Body).Decode(&newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := validUser(newUser); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		uniq, err := h.store.CheckUserByNameEmail(newUser.Email, newUser.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if uniq {
			http.Error(w, " Username or Email is already in used! ", http.StatusBadRequest)
			return
		}
		hashedPassword, err := store.GenerateHashPassword(newUser.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		newUser.Password = string(hashedPassword)
		if err := h.store.CreateUser(newUser); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "You have successfully registered ")

	case http.MethodGet:
		fmt.Fprintf(w, "Welcome to register page")
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
