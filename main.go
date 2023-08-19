package main

import (
	"fmt"
	"log"
	"net/http"
	"restful-api/pkg"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	env, _ := pkg.LoadConfig()

	_ = pkg.DbInit(env)

	//routes.RouteInit(router, db, env)

	log.Printf("Server berjalan di http://localhost:%s", env.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", env.Port), router))
}
