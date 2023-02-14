package example

import (
	"GOLANG-DDD-MONGO-TEMPLATE/shared"
	"GOLANG-DDD-MONGO-TEMPLATE/shared/dto/example"
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
