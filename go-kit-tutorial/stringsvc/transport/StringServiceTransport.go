package transport

import (
	. "golang.org/x/net/context"

	"encoding/json"
	. "github.com/1ambda/golang/go-kit-tutorial/stringsvc/service"
	"github.com/pkg/errors"
	"net/http"
)

func DecodeUppercaseRequest(_ Context, r *http.Request) (interface{}, error) {
	var request UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "Failed to decode UppercaseRequest")
	}
	return request, nil
}

func DecodeCountRequest(_ Context, r *http.Request) (interface{}, error) {
	var request CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, errors.Wrap(err, "Failed to decode CountRequest")
	}
	return request, nil
}

func EncodeResponse(_ Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
