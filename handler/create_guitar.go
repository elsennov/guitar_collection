package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"personal/guitar_collection/domain"
	"personal/guitar_collection/service"
)

func CreateGuitarHandler(services *service.Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var viewGuitar domain.ViewGuitar
		err := json.NewDecoder(r.Body).Decode(&viewGuitar)
		defer r.Body.Close()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = services.Guitar.Process(viewGuitar)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
