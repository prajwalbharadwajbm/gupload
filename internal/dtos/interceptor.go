package dtos

type InterceptorResponse struct {
	Data         any    `json:"data"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
}
