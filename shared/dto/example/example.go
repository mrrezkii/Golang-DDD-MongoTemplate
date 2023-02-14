package example

import "GOLANG-DDD-MONGO-TEMPLATE/shared/dto"

type (
	ExampleRequestDto struct {
		MandatoryRequest dto.MandatoryRequestDto
	}

	ExampleResponseDto struct {
		RequestID string `json:"request_id"`
	}
)
