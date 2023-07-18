package main

import (
	"fmt"
	"github-dbq/config"
	routes "github-dbq/src/student/http"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	env, _ := config.LoadConfig()

	db := config.DbInit(env)

	routes.RouteInit(router, db, env)

	log.Printf("Server berjalan di http://localhost:%s", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.Port), router))
}
