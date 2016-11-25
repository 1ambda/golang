package service

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
)

func createUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}

		return UppercaseResponse{v, ""}, nil
	}
}

func CreateUppercaseHandler(ctx context.Context, svc StringService) *http.Server {
	return http.NewServer(
		ctx,
		createUppercaseEndpoint(svc),
		DecodeUppercaseRequest,
		EncodeResponse,
	)
}

func createCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}

func CreateCountHandler(ctx context.Context, svc StringService) *http.Server {
	return http.NewServer(
		ctx,
		createCountEndpoint(svc),
		DecodeCountRequest,
		EncodeResponse,
	)
}
