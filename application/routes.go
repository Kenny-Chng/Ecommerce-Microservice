package application 

import (
	"net/http"
	"github.com/go-chi/chi/v5 "
)

func loadRoutes () *chi.Mux {
	router := chi.newRouter()
	router.Use(middleware.Logger)
	router.Get("/", func (w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	} )
	return router
}