package routes

import (
	nationalityController "be-family/controllers/nationality"

	"github.com/gorilla/mux"
)

func NationalityRoutes(r *mux.Router) {
	router := r.PathPrefix("/nationality").Subrouter()

	router.HandleFunc("", nationalityController.Index).Methods("GET")
	router.HandleFunc("", nationalityController.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", nationalityController.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", nationalityController.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", nationalityController.Delete).Methods("DELETE")
}
