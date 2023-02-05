package example

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/application"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto/example"
	"context"
)

type (
	ViewService interface {
		FindAll(ctx context.Context, request example.ExampleRequestDto) (*example.ExampleResponseDto, error)
	}

	service struct {
		SharedHolder      shared.Holder
		ApplicationHolder application.Holder
	}
)

func NewViewService(sh shared.Holder, ah application.Holder) (ViewService, error) {
	return &service{sh, ah}, nil
}
