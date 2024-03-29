package responses

import (
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
)

func GetCustomerResponse(customer *models.Customer) *CommonResponse {
	return &CommonResponse{
		Data: customer,
	}
}

func GetCustomersResponse(customer []models.Customer, count *int, pagination requests.PaginationRequest) *CommonResponse {
	itemsInPage := len(customer)
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
		Data: customer,
	}
}
