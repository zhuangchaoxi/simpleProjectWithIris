package response

type Response struct {
	Code        int         `json:"code"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	ErrorDetail string      `json:"error_detail"`
}

func HttpResponse(code int, msg string, data interface{}, errDetail string) *Response {
	return &Response{code, msg, data, errDetail}
}
