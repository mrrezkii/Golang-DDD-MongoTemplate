package example

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/application"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto/example"
	"context"
)

type (
	Service interface {
		FindAll(ctx context.Context, request example.ExampleRequestDto) (*example.ExampleResponseDto, error)
	}

	ServiceImpl struct {
		Service           *Service
		ApplicationHolder application.Holder
	}
)