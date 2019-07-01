package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"personal/guitar_collection/handler"
	"personal/guitar_collection/service"
)

func Router(services *service.Services) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/v1/guitar", logging(handler.CreateGuitarHandler(services))).Methods("POST")
	return router
}

func logging(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handlerFunc(writer, request)
	}
}
