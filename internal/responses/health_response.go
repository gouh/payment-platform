package responses

func GetHealthResponse() *CommonResponse {
	return &CommonResponse{
		Data: "Ok",
	}
}
