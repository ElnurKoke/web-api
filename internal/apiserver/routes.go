package apiserver

import (
	"net/http"
)

func (s *Server) InitRoutes() http.Handler {
	s.router.HandleFunc("/", s.home)
	s.router.HandleFunc("/register", s.register)
	s.router.HandleFunc("/login", s.login)
	s.router.HandleFunc("/logout", s.logOut)
	s.router.HandleFunc("/update/name", s.middleWareGetUser(s.updateName))
	s.router.HandleFunc("/update/email", s.middleWareGetUser(s.updateEmail))
	s.router.HandleFunc("/update/project/", s.middleWareGetUser(s.updateProjectByAdmin))
	return s.router
}
