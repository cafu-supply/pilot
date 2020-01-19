package repository

import "../domain/entity"

type PilotRepo interface {
	ListPilots() ([]entity.Pilot, error)
	UpdatePilotStatus(*entity.Pilot) (*entity.Pilot, error)
	//GetPilot(id string) (Pilot, error)
	//CreatePilot(pilot-management CreatePilotParams) (Pilot, error)
	//UpdatePilot(pilot-management UpdatePilotParams) (Pilot, error)
	//DeletePilot(id string) error
}
