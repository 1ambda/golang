package endpoint

import (
	. "github.com/go-kit/kit/endpoint"
	. "golang.org/x/net/context"

	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc2/service"
)

func MakeUppercaseEndpoint(svc StringService) Endpoint {
	return func(ctx Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}

		return UppercaseResponse{v, ""}, nil
	}
}

func MakeCountEndpoint(svc StringService) Endpoint {
	return func(ctx Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(req.S)
		return CountResponse{v}, nil
	}
}
