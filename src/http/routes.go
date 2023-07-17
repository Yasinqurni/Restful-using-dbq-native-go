package routes

import (
	"database/sql"
	"github-dbq/config"
	"github-dbq/src"

	"github.com/gorilla/mux"
)

func RouteInit(router *mux.Router, db *sql.DB, env *config.Config) {

	studentController, _, _ := src.InitModule(db, env)

	router.HandleFunc("/student/{id}", studentController.GetByID).Methods("GET")
	router.HandleFunc("/student", studentController.Get).Methods("GET")

}
