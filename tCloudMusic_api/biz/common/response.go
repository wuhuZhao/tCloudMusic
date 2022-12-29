package common

type Response struct {
	data interface{}
	msg  string
}

func NewErrorResponse(msg string) *Response {
	return &Response{data: "", msg: msg}
}

func NewSuccessResponse(data interface{}) *Response {
	return &Response{data: data, msg: ""}
}
