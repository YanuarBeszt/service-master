package router

import (
	"go-postgres-crud/controller"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/employee", controller.GetAllEmployee).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/employee/{id}", controller.GetEmployee).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/employee", controller.addEmployee).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/employee/{id}", controller.GetAllEmployee).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/employee/{id}", controller.GetAllEmployee).Methods("DELETE", "OPTIONS")

	return router
}
