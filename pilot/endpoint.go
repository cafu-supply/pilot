package pilot

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeListPilotsEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := s.ListPilots()
		if err != nil {
			return Response{data: nil, errors: []error{err}}, err
		}
		var pilots []PilotView
		for _, pilot := range v {
			pilots = append(pilots, toPilotView(pilot))
		}
		return Response{data: pilots, errors: nil}, nil
	}
}

func makeGetPilot(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPilotRequest)
		pilot, err := s.GetPilot(req.Id)
		if err != nil {
			return Response{data: nil, errors: []error{err}}, err
		}
		return Response{data: toPilotView(pilot), errors: nil}, nil
	}
}

func makeCreatePilot(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePilotRequest)
		pilot, err := s.CreatePilot(CreatePilotParams(req))
		if err != nil {
			return Response{data: nil, errors: []error{err}}, err
		}
		return Response{data: toPilotView(pilot), errors: nil}, nil
	}
}

func makeUpdatePilot(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePilotRequest)
		pilot, err := s.UpdatePilot(UpdatePilotParams(req))
		if err != nil {
			return Response{data: nil, errors: []error{err}}, err
		}
		return Response{data: toPilotView(pilot), errors: nil}, nil
	}
}

func makeDeletePilot(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(DeletePilotRequest)
		err := s.DeletePilot(req.Id)
		if err != nil {
			return Response{data: nil, errors: []error{err}}, err
		}
		return Response{data: nil, errors: nil}, nil
	}
}
