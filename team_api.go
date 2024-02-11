package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func teamHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
	case http.MethodGet:

		vars := mux.Vars(r)

		id := vars[ID]

		log.Println(id)

		j, err := json.Marshal(id)

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
	
} // teamHandler
