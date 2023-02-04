package util

import (
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/dto"
	"SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/util/constant"
	tasha_error "SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/util/error"
	"github.com/labstack/echo"
	"net/http"
	"strings"
	"time"
)

func GetMandatoryRequest(ec echo.Context) dto.MandatoryRequestDto {
	var (
		header = ec.Request().Header
	)

	return dto.MandatoryRequestDto{
		StoreID:   getHeaderV2OrDefault(header, "X-Store-Id", ""),
		ChannelID: getHeaderV2OrDefault(header, "X-Channel-Id", ""),
		RequestID: getHeaderV2OrDefault(header, "X-Request-Id", ""),
		Username:  getHeaderV2OrDefault(header, "X-Username", ""),
		Language:  getLanguage(header),
	}
}

func Response(ec echo.Context, data interface{}, err error) error {
	if err == nil {
		successResponse := tasha_error.New(tasha_error.SUCCESS, nil)
		return ec.JSON(successResponse.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       successResponse.GetCode(),
			Message:    successResponse.GetMessage(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
	if s, ok := err.(tasha_error.ErrorStandard); ok {
		return ec.JSON(s.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       s.GetCode(),
			Message:    s.GetMessage(),
			Errors:     s.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	} else {
		errResponse := tasha_error.New(tasha_error.SYSTEM_ERROR, err)
		return ec.JSON(errResponse.GetHTTPStatus(), &dto.BaseResponseDto{
			Code:       errResponse.GetCode(),
			Message:    errResponse.GetMessage(),
			Errors:     errResponse.GetErrors(),
			Data:       data,
			ServerTime: time.Now().Unix(),
		})
	}
}

func getLanguage(header http.Header) string {
	language := getHeaderV2OrDefault(header, "Accept-Language", "")
	language = strings.ToLower(language)
	if strings.HasPrefix(language, constant.LangEn) {
		return constant.LangEn
	}
	if strings.HasPrefix(language, constant.LangId) {
		return constant.LangId
	}
	return constant.LangEn
}

func getHeaderV2OrDefault(header http.Header, key, value string) string {
	var (
		val = header.Get(key)
	)

	if len(val) == 0 {
		return value
	}

	return val
}
