package endpoint

import (
	"../domain"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeListPilotsEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		v, err := s.ListPilots()
		if err != nil {
			return Response{Data: nil, Errors: []error{err}}, err
		}
		pilots := make([]PilotView, 0)
		for _, pilot := range v {
			pilots = append(pilots, toPilotView(pilot))
		}
		return Response{Data: pilots, Errors: nil}, nil
	}
}

func MakeUpdatePilotStatusEndpoint(s domain.Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdatePilotStatusRequest)
		par := domain.UpdatePilotStatusParams(req)
		v, err := s.UpdatePilotStatus(&par)
		if err != nil {
						return Response{Data: nil, Errors: []error{err}}, err
					}
					return Response{Data: toPilotView(*v), Errors: nil}, nil
	}
}

//
//func MakeGetPilot(s Service) endpoint.Endpoint {
//	return func(_ context.Context, request interface{}) (interface{}, error) {
//		req := request.(GetPilotRequest)
//		pilot-management, err := s.GetPilot(req.Id)
//		if err != nil {
//			return Response{data: nil, errors: []error{err}}, err
//		}
//		return Response{data: toPilotView(pilot-management), errors: nil}, nil
//	}
//}
//
//func MakeCreatePilot(s Service) endpoint.Endpoint {
//	return func(_ context.Context, request interface{}) (interface{}, error) {
//		req := request.(CreatePilotRequest)
//		pilot-management, err := s.CreatePilot(CreatePilotParams(req))
//		if err != nil {
//			return Response{data: nil, errors: []error{err}}, err
//		}
//		return Response{data: toPilotView(pilot-management), errors: nil}, nil
//	}
//}
//
//func MakeUpdatePilot(s Service) endpoint.Endpoint {
//	return func(_ context.Context, request interface{}) (interface{}, error) {
//		req := request.(UpdatePilotRequest)
//		pilot-management, err := s.UpdatePilot(UpdatePilotParams(req))
//		if err != nil {
//			return Response{data: nil, errors: []error{err}}, err
//		}
//		return Response{data: toPilotView(pilot-management), errors: nil}, nil
//	}
//}
//
//func MakeDeletePilot(s Service) endpoint.Endpoint {
//	return func(_ context.Context, request interface{}) (interface{}, error) {
//		req := request.(DeletePilotRequest)
//		err := s.DeletePilot(req.Id)
//		if err != nil {
//			return Response{data: nil, errors: []error{err}}, err
//		}
//		return Response{data: nil, errors: nil}, nil
//	}
//}
