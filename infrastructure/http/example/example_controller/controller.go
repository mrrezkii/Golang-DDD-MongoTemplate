package example_controller

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared"
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
