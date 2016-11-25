package service

import (
	"strings"

	"github.com/pkg/errors"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringMiddleware func(StringService) StringService

type StringServiceImpl struct{}

func (StringServiceImpl) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("Empty string")
	}

	return strings.ToUpper(s), nil
}

func (StringServiceImpl) Count(s string) int {
	return len(s)
}

type UppercaseRequest struct {
	S string `json:"s"`
}

type UppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type CountRequest struct {
	S string `json:"s"`
}

type CountResponse struct {
	V int `json:"v"`
}
