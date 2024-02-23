package requests

import (
	"braces.dev/errtrace"
	"github.com/gorilla/schema"
)

type PaginationRequest struct {
	Page     int `schema:"page"`
	PageSize int `schema:"size"`
}

func (paginationRequest *PaginationRequest) QueryParamsToStruct(
	queryValues map[string][]string, queryParams interface{},
) error {
	decoder := schema.NewDecoder()

	paginationRequest.Page = 1
	paginationRequest.PageSize = 10

	err := decoder.Decode(queryParams, queryValues)
	if err != nil {
		return errtrace.Wrap(err)
	}
	return nil
}
