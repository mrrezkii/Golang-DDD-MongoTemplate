package dto

import "time"

type (
	MandatoryRequestDto struct {
		StoreID   string `json:"X-Store-ID" validate:"required"`
		ChannelID string `json:"X-Channel-ID" validate:"required"`
		RequestID string `json:"X-Request-ID" validate:"required"`
		Username  string `json:"X-Username" validate:"required"`
		Language  string `json:"Accept-Language" validate:"required"`
	}

	BaseResponseDto struct {
		Code       string      `json:"code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Errors     []string    `json:"errors"`
		ServerTime int64       `json:"server_time"`
	}

	BaseTableFields struct {
		ID          uint64    `json:"id"`
		CreatedDate time.Time `json:"createdDate"`
		UpdatedDate time.Time `json:"updatedDate"`
		CreatedBy   string    `json:"createdBy"`
		UpdatedBy   string    `json:"updatedBy"`
		Version     int       `json:"version"`
		IsDeleted   bool      `json:"isDeleted"`
	}
)
