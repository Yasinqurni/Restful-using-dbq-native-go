package routes

import (
	"database/sql"
	"github-dbq/config"
	src "github-dbq/src/student"

	"github.com/gorilla/mux"
)

func RouteInit(router *mux.Router, db *sql.DB, env *config.Config) {

	studentController, _, _ := src.InitModule(db, env)

	router.HandleFunc("/student/{id}", studentController.GetByID).Methods("GET")
	router.HandleFunc("/student", studentController.Get).Methods("GET")
	router.HandleFunc("/student/{id}", studentController.Update).Methods("PUT")
	router.HandleFunc("/student/{id}", studentController.Delete).Methods("DELETE")
	router.HandleFunc("/student", studentController.Create).Methods("POST")

}
