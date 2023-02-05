package main

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/infrastructure"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/interfaces"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared"
	"github.com/go-playground/validator"
)

func main() {
	sh := shared.Holder{
		Config:    nil,
		Echo:      nil,
		Logger:    nil,
		Validator: validator.Validate{},
	}
	inh := infrastructure.Holder{
		ExampleController: nil,
		SharedHolder:      shared.Holder{},
		InterfaceHolder:   interfaces.Holder{},
	}

	go inh.ListenHTTP()

	sh.Close()
}
