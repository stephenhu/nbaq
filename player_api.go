package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func playerHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
	case http.MethodGet:

		vars := mux.Vars(r)

		id := vars[ID]

		log.Println(id)

		g := cache.Seasons["2023"].Players[1626167]

		j, err := json.Marshal(g)

		if err != nil {
			log.Println(err)
		} else {
			w.Write(j)
		}

	case http.MethodDelete:
	case http.MethodPost:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	
} // playerHandler
