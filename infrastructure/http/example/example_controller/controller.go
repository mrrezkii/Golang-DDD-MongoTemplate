package example_controller

import (
	"GOLANG-DDD-MONGO-TEMPLATE/interfaces"
	"GOLANG-DDD-MONGO-TEMPLATE/shared"
)

type (
	Controller struct {
		SharedHolder    shared.Holder
		InterfaceHolder interfaces.Holder
	}
)

func NewController(sh shared.Holder, ih interfaces.Holder) (*Controller, error) {
	return &Controller{sh, ih}, nil
}
