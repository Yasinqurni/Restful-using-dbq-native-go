package routes

import (
	"database/sql"
	"github-dbq/src"

	"github.com/gorilla/mux"
)

func RouteInit(router *mux.Router, db *sql.DB) {

	studentController, _, _ := src.InitModule(db)

	router.HandleFunc("/student/{id}", studentController.GetByID).Methods("GET")
	router.HandleFunc("/student", studentController.Get).Methods("GET")

}
