package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


var (
  src 			string
	cache			NbaCache
)


func initFlags() {
  flag.StringVar(&src, "src", CURRENT_DIR, "source directory")
} // initFlags


func initRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/games/{id:[0-9]+}", gameHandler)
	router.HandleFunc("/api/v1/players/{id:[0-9]+}", playerHandler)
	router.HandleFunc("/api/v1/teams/{id:[0-9]+}", teamHandler)
	router.HandleFunc("/api/v1/version", versionHandler)

	return router

} // initRouter


func main() {

	initFlags()

	flag.Parse()

	fmt.Printf("Starting %s...\n", version())

	initCache()

	log.Fatal(http.ListenAndServe("127.0.0.1:8000", initRouter()))

} // main
