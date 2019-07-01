package server

import (
	"net/http"
	"personal/guitar_collection/service"
)

func StartAPIServer() {
	router := Router(service.NewServices())
	err := http.ListenAndServe(":80", router)
	if err != nil {
		panic(err)
	}
}
