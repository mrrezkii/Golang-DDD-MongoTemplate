package infrastructure

import (
	"GOLANG-DDD-MONGO-TEMPLATE/shared/util"
	utilctx "GOLANG-DDD-MONGO-TEMPLATE/shared/util/context"
	tasha_error "GOLANG-DDD-MONGO-TEMPLATE/shared/util/error"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"strings"
)

const (
	HeaderXRequestID = "X-Request-Id"
)

// RegisterMiddleware registers middlewares used in all endpoint
func RegisterMiddleware(holder *Holder) {
	holder.SharedHolder.Echo.Use(SetMandatoryRequest())
}

func SetMandatoryRequest() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mandatory := util.GetMandatoryRequest(c)
			err := c.Validate(&mandatory)
			if err != nil && isApplySetMandatoryRequestChecking(c.Request().RequestURI) {
				return util.Response(c, nil, tasha_error.New(tasha_error.BAD_REQUEST, err))
			}

			// set requestID
			if mandatory.RequestID == "ee35b46d" || mandatory.RequestID == "" {
				mandatory.RequestID = uuid.New().String()
			}

			ctx := utilctx.SetMandatoryRequest(c.Request().Context(), mandatory)
			c.SetRequest(c.Request().WithContext(ctx))

			// set requestID in http response
			c.Response().Header().Set(HeaderXRequestID, mandatory.RequestID)

			return next(c)
		}
	}
}

func isApplySetMandatoryRequestChecking(uri string) bool {
	if strings.HasPrefix(uri, "/application/health") || strings.HasPrefix(uri, "/swagger") {
		return false
	}
	return true
}
