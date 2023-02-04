package context

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto"
	"context"
)

var contextKeyMandatoryRequest = contextKey("MandatoryRequest")

func MandatoryRequest(c context.Context) dto.MandatoryRequestDto {
	mandatoryRequest := c.Value(contextKeyMandatoryRequest)
	if mandatoryRequest == nil {
		return dto.MandatoryRequestDto{}
	}
	return mandatoryRequest.(dto.MandatoryRequestDto)
}

func SetMandatoryRequest(c context.Context, mandatoryRequest dto.MandatoryRequestDto) context.Context {
	return context.WithValue(c, contextKeyMandatoryRequest, mandatoryRequest)
}
