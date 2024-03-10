package apiserver

import (
	"log"
	"net/http"
	"time"

	"github.com/ElnurKoke/web-api.git/internal/model"
	"github.com/ElnurKoke/web-api.git/internal/store"
)

//start
func Start(config *model.Config) error {
	db, err := store.NewDB(config.DB.Dsn, config.DB.Driver)
	if err != nil {
		return err
	}
	defer db.Close()

	store := store.New(db)
	srv := newServer(store)

	httpserver := &http.Server{
		Addr:           config.Port,
		Handler:        srv.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	log.Printf("Server run on http://localhost%s", config.Port)
	return httpserver.ListenAndServe()
}
