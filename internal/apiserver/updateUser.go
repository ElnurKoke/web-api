package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ElnurKoke/web-api.git/internal/model"
)

func (s *Server) updateName(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/update/name" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	user, ok := r.Context().Value("user").(model.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if !user.IsAuth {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validName(newUser.Username); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exist, err := s.store.CheckUserByName(newUser.Username)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if exist {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}
	if err := s.store.UpdateUserName(user.Id, newUser.Username); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You have successfully update name")
}

func (s *Server) updateEmail(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/update/email" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	user, ok := r.Context().Value("user").(model.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if !user.IsAuth {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validEmail(newUser.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	exist, err := s.store.CheckUserByEmail(newUser.Email)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if exist {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}
	if err := s.store.UpdateUserEmail(user.Id, newUser.Email); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "You have successfully update email")
}
