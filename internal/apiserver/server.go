package apiserver

import (
	"net/http"

	"github.com/ElnurKoke/web-api.git/internal/store"
)

type Server struct {
	router *http.ServeMux
	store  *store.Store
}

func newServer(store *store.Store) *Server {
	return &Server{
		router: http.NewServeMux(),
		store:  store,
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
