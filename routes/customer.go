package routes

import (
	customerController "be-family/controllers/customer"

	"github.com/gorilla/mux"
)

func CustomerRoutes(r *mux.Router) {
	router := r.PathPrefix("/customer").Subrouter()

	router.HandleFunc("", customerController.Index).Methods("GET")
	router.HandleFunc("", customerController.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", customerController.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", customerController.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", customerController.Delete).Methods("DELETE")
}
