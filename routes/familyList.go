package routes

import (
	costumerController "be-family/controllers/familyList"

	"github.com/gorilla/mux"
)

func FamilyListRoutes(r *mux.Router) {
	router := r.PathPrefix("/familylist").Subrouter()

	router.HandleFunc("", costumerController.Index).Methods("GET")
	router.HandleFunc("", costumerController.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", costumerController.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", costumerController.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", costumerController.Delete).Methods("DELETE")
}
