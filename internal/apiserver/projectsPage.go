package apiserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ElnurKoke/web-api.git/internal/model"
)

func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	projects, err := s.store.GetAllProjects()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Welcome to main page, projects:  ")
		json.NewEncoder(w).Encode(projects)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (s *Server) updateProjectByAdmin(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/update/project/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 || err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	user, ok := r.Context().Value("user").(model.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if user.Role != "admin" {
		http.Error(w, "need admin role", http.StatusUnauthorized)
		return
	}
	project, err := s.store.GetProjectById(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodPut:
		var newProject model.Project
		err := json.NewDecoder(r.Body).Decode(&newProject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newProject = mergeProjects(project, newProject)
		newProject.ID = id
		if err := s.store.UpdateProject(newProject); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Project updated ")
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
