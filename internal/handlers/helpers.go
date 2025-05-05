package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, data interface{}, status int, err error) {
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Println(err)
	}

}
