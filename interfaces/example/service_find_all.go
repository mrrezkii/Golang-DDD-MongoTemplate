package example

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto/example"
	"context"
	"github.com/pkg/errors"
)

func (s *service) FindAll(ctx context.Context, request example.ExampleRequestDto) (*example.ExampleResponseDto, error) {
	res, err := s.ApplicationHolder.ExampleService.FindAll(ctx, request)

	if err != nil {
		return nil, errors.Wrap(err, "failed to proceed FindAll")
	}
	return res, nil
}
