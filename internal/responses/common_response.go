package responses

type Meta struct {
	Page        *int    `json:"page"`
	PageSize    *int    `json:"size"`
	ItemsInPage *int    `json:"itemsInPage"`
	TotalPages  *int    `json:"totalPages"`
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
