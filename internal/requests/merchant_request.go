package requests

type MerchantParams struct {
	PaginationRequest
	Id   *string `schema:"id"`
	Name *string `schema:"name"`
}
