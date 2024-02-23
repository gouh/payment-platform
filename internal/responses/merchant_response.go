package responses

import (
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
)

func GetMerchantResponse(merchant *models.Merchant) *CommonResponse {
	return &CommonResponse{
		Data: merchant,
	}
}

func GetMerchantsResponse(merchants []models.Merchant, count *int, pagination requests.PaginationRequest) *CommonResponse {
	itemsInPage := len(merchants)
	totalPages := 0

	if pagination.PageSize > 0 {
		totalPages = (*count + pagination.PageSize - 1) / pagination.PageSize
	}

	return &CommonResponse{
		Meta: &Meta{
			Page:        &pagination.Page,
			PageSize:    &pagination.PageSize,
			ItemsInPage: &itemsInPage,
			TotalPages:  &totalPages,
			Error:       nil,
		},
		Data: merchants,
	}
}
