package pilot

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

type UpdatePilotRequest struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	CodeName   string `json:"codeName"`
	SupplierId string `json:"supplierId"`
	MarketId   string `json:"marketId"`
	ServiceId  string `json:"serviceId"`
}
