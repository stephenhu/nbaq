package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/madsportslab/sportsearch"
)


func searchHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPut:
	case http.MethodGet:

		q := r.URL.Query().Get("q")

		log.Println(q)

		sportsearch.InitMaps()

		res := sportsearch.Classifier(q)

		log.Println(res)

		j, err := json.Marshal("{q: abc}")

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
	
} // searchHandler
