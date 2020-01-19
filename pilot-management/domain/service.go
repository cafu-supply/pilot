package domain

import "./entity"

type Service interface {
	ListPilots() ([]entity.Pilot, error)
	UpdatePilotStatus(params *UpdatePilotStatusParams) (*entity.Pilot, error)
	//GetPilot(id string) (Pilot, error)
	//CreatePilot(params CreatePilotParams) (Pilot, error)
	//UpdatePilot(params UpdatePilotParams) (Pilot, error)
	//DeletePilot(id string) error
}

type CreatePilotParams struct {
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
}

type UpdatePilotParams struct {
	Id         string
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
}

type UpdatePilotStatusParams struct {
	Id         string
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
	Status     string
}