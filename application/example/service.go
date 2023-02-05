package example

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto/example"
	"context"
)

type (
	Service interface {
		FindAll(ctx context.Context, request example.ExampleRequestDto) (*example.ExampleResponseDto, error)
	}

	service struct {
		SharedHolder shared.Holder
	}
)

func NewService(holder shared.Holder) (Service, error) {
	return &service{holder}, nil
}
