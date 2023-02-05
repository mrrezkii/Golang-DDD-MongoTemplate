package example

import "SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto"

type (
	ExampleRequestDto struct {
		MandatoryRequest dto.MandatoryRequestDto
	}

	ExampleResponseDto struct {
		RequestID string `json:"request_id"`
	}
)
