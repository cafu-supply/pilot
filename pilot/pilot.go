package pilot

// model.go
type Pilot struct {
	Id         string
	UserId     string
	CodeName   string
	SupplierId string
	MarketId   string
	ServiceId  string
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

// service.go
type Service interface {
	ListPilots() ([]Pilot, error)
	GetPilot(id string) (Pilot, error)
	CreatePilot(params CreatePilotParams) (Pilot, error)
	UpdatePilot(params UpdatePilotParams) (Pilot, error)
	DeletePilot(id string) error
}

// repo.go
type PilotRepo interface {
	ListPilots() ([]Pilot, error)
	GetPilot(id string) (Pilot, error)
	CreatePilot(pilot Pilot) (Pilot, error)
	UpdatePilot(pilot Pilot) (Pilot, error)
	DeletePilot(id string) error
}

type ServiceImpl struct {
	pilotRepo PilotRepo
}

func (s ServiceImpl) ListPilots() ([]Pilot, error) {
	return s.pilotRepo.ListPilots()
}

func (s ServiceImpl) GetPilot(id string) (Pilot, error) {
	return s.pilotRepo.GetPilot(id)
}

func (s ServiceImpl) CreatePilot(pilot Pilot) (Pilot, error) {
	return s.pilotRepo.CreatePilot(pilot)
}

func (s ServiceImpl) UpdatePilot(pilot Pilot) (Pilot, error) {
	return s.pilotRepo.UpdatePilot(pilot)
}

func (s ServiceImpl) DeletePilot(id string) error {
	return s.pilotRepo.DeletePilot(id)
}
