package pilot-management

// model.go
type Pilot struct {
	Id   string
	UserId string
	CodeName string
	SupplierId string
	MarketId string
	ServiceId string
}

// servcie.go
type Service interface {
	ListPilots() ([]Pilot, error)
	GetPilot(id string) (Pilot, error)
	CreatePilot(pilot Pilot) (Pilot, error)
	UpdatePilot(pilot Pilot) (Pilot, error)
	DeletePilot(id string) (interface{}, error)
}

// repo.go
type PilotRepo interface {
	ListPilots() ([]Pilot, error)
	GetPilot(id string) (Pilot, error)
	CreatePilot(pilot Pilot) (Pilot, error)
	UpdatePilot(pilot Pilot) (Pilot, error)
	DeletePilot(id string) (interface{}, error)
}

type service struct {
	pilotRepo PilotRepo
}

func (s service) ListPilots() ([]Pilot, error) {
	return s.pilotRepo.ListPilots()
}

func (s service) GetPilot(id int64) (Pilot, error) {
	return s.pilotRepo.GetPilot(id)
}

func (s service) CreatePilot(name string) (Pilot, error) {
	return s.pilotRepo.CreatePilot(name)
}

func (s service) UpdatePilot(Pilot Pilot) (Pilot, error) {
	return s.pilotRepo.UpdatePilot(Pilot)
}

func (s service) DeletePilot(id int64) (interface{}, error) {
	return s.pilotRepo.DeletePilot(id)
}

