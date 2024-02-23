package requests

type CustomerParams struct {
	PaginationRequest
	Id    *string `schema:"id"`
	Name  *string `schema:"name"`
	Email *string `schema:"email"`
}
