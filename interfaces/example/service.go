package example

import (
	"GOLANG-DDD-MONGO-TEMPLATE/application"
	"GOLANG-DDD-MONGO-TEMPLATE/shared"
	"GOLANG-DDD-MONGO-TEMPLATE/shared/dto/example"
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
