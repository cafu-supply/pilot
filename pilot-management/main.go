package main

import (
	"./service"
	"./endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func assignRoutes(router *mux.Router) *mux.Router {
	service := service.MakeServiceImpl()

	listPilotsHandler := httpTransport.NewServer(
		endpoint.MakeListPilotsEndpoint(service),
		endpoint.MakeDecoder(endpoint.ListPilotsRequest{}),
		endpoint.EncodeResponse)

	updatePilotStatusHandler := httpTransport.NewServer(
		endpoint.MakeUpdatePilotStatusEndpoint(service),
		endpoint.MakeDecoder(endpoint.UpdatePilotStatusRequest{}),
		endpoint.EncodeResponse)

	router.Handle("/pilots", listPilotsHandler).Methods("GET")
	router.Handle("/pilots", updatePilotStatusHandler).Methods("PUT")

	return router
}

func main() {
	router := mux.NewRouter()
	assignRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
