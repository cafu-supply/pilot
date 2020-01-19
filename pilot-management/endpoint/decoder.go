package endpoint

import (
	"context"
	"encoding/json"
	"net/http"
)

type ListPilotsRequest struct{}

type GetPilotRequest struct {
	Id string `json:"id"`
}

type DeletePilotRequest struct {
	Id string `json:"id"`
}

type CreatePilotRequest struct {
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
}

//type UpdatePilotRequest struct {
//	Id         string `json:"id"`
//	UserId     string `json:"userId"`
//	CodeName   string `json:"codeName"`
//	SupplierId string `json:"supplierId"`
//	MarketId   string `json:"marketId"`
//	ServiceId  string `json:"serviceId"`
//}

type UpdatePilotStatusRequest struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
	Status     string `json:"status"`
}

//func DecodeListPilotsRequest(_ context.Context, r *http.Request) (interface{}, error) {
//	var request ListPilotsRequest
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		return nil, err
//	}
//	return request, nil
//}

func MakeDecoder(request interface{}) func (_ context.Context, r *http.Request) (interface{}, error) {
	return func (_ context.Context, r *http.Request) (interface{}, error) {
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			return nil, err
		}
		return request, nil
	}
}
