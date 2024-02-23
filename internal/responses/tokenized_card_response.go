package responses

import (
	"payment-platform/internal/models"
	"payment-platform/internal/requests"
)

func GetTokenizedCardResponse(tokenizedCard *models.TokenizedCard) *CommonResponse {
	return &CommonResponse{
		Data: tokenizedCard,
	}
}

func GetTokenizedCardsResponse(tokenizedCards []models.TokenizedCard, count *int, pagination requests.PaginationRequest) *CommonResponse {
	itemsInPage := len(tokenizedCards)
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
		Data: tokenizedCards,
	}
}
