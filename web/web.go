package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"supply/pilot"
)

func assignRoutes(router mux.Router) mux.Router {
	service := pilot.ServiceImpl{}

	router.HandleFunc("/pilots", httptransport.NewServer())
}
