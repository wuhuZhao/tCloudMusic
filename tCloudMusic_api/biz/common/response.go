package common

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func NewErrorResponse(msg string) Response {
	return Response{Data: "", Msg: msg}
}

func NewSuccessResponse(data interface{}) Response {
	return Response{Data: data, Msg: ""}
}
