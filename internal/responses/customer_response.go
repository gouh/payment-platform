package responses

import (
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
)

func GetCustomerResponse(merchant *models.Customer) *CommonResponse {
	return &CommonResponse{
		Data: merchant,
	}
}

func GetCustomersResponse(merchants []models.Customer, count *int, pagination requests.PaginationRequest) *CommonResponse {
	itemsInPage := len(merchants)
	totalPages := 1
	if *count >= pagination.PageSize {
		totalPages = *count / pagination.PageSize
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
