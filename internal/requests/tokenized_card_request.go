package requests

type TokenizedCardParams struct {
	PaginationRequest
	Token      *string `schema:"token"`
	CustomerId *string `schema:"customerId"`
}

type TokenizedCardRequest struct {
	CardNumber  string `json:"cardNumber" binding:"required"`
	ExpiryMonth int    `json:"expiryMonth" binding:"required"`
	ExpiryYear  int    `json:"expiryYear" binding:"required"`
	CardType    string `json:"cardType" binding:"required"`
}
