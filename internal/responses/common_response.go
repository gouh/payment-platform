package responses

type Meta struct {
	Page        *int    `json:"page,omitempty"`
	PageSize    *int    `json:"size,omitempty"`
	ItemsInPage *int    `json:"itemsInPage,omitempty"`
	TotalPages  *int    `json:"totalPages,omitempty"`
	Error       *string `json:"error,omitempty"`
}

type CommonResponse struct {
	Meta *Meta       `json:"meta,omitempty"`
	Data interface{} `json:"data"`
}

func GetErrorResponse(err string) *CommonResponse {
	return &CommonResponse{
		Meta: &Meta{
			Error: &err,
		},
	}
}
