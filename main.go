package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


var (
  dir 			string
)


func initFlags() {
  flag.StringVar(&dir, "dir", CURRENT_DIR, "source directory")
} // initFlags


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/games/{id:[0-9]+}", gameHandler)
	router.HandleFunc("/api/v1/players/{id:[0-9]+}", playerHandler)
	router.HandleFunc("/api/v1/search", searchHandler)
	router.HandleFunc("/api/v1/teams/{id:[0-9]+}", teamHandler)
	router.HandleFunc("/api/v1/version", versionHandler)

	return router

} // initRouter


func main() {

	fmt.Printf("Starting %s...\n", version())

	initWarehouse()

	defer db.Close()

	log.Fatal(http.ListenAndServe("localhost:8000", initRouter()))

} // main
