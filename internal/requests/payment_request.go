package requests

type PaymentParams struct {
	PaginationRequest
	Id            *string `schema:"id"`
	MerchantId    *string `schema:"merchantId"`
	CustomerId    *string `schema:"customerId"`
	TokenizedCard *string `schema:"tokenizedCard"`
}
