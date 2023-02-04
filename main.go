package main

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/log"
	utilctx "SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/util/context"
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		l   = logrus.New()
		ctx = utilctx.SetMandatoryRequest(context.Background(), dto.MandatoryRequestDto{
			RequestID: uuid.New().String(),
			Username:  "sandbox",
		})
	)

	logger, _ := log.NewLogger(l)
	logger.Info(ctx, "Listening on port 7001")
}
