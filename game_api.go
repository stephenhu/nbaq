package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func gameHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
	case http.MethodGet:

		vars := mux.Vars(r)

		id := vars[ID]

		log.Println(id)

		g := cache.Seasons["2023"].Games[id]

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
	
} // gameHandler
