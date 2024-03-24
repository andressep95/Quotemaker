package response

type ResponseData struct {
	Result interface{} `json:"result"`
}
type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type Response struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"status_code"`
	Message    string        `json:"message"`
	ErrorCode  string        `json:"error_code,omitempty"`
	Errors     []ErrorDetail `json:"errors,omitempty"`
	Data       ResponseData  `json:"data,omitempty"`
}
