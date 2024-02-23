package responses

import (
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
)

func GetPaymentResponse(payment *models.Payment) *CommonResponse {
	return &CommonResponse{
		Data: payment,
	}
}

func GetPaymentsResponse(payments []models.Payment, count *int, pagination requests.PaginationRequest) *CommonResponse {
	itemsInPage := len(payments)
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
		Data: payments,
	}
}
