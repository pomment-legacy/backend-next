package utils

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Content interface{} `json:"content"`
}

const (
	MsgBadArgument = "Bad Argument"
	MsgBadToken = "Bad Token"
	MsgGeneralFailure = "Server Failure"
	MsgSuccess = "Success"
)

func FailureRes(msg string) Response {
	return Response{
		Success: false,
		Message: msg,
		Content: nil,
	}
}

func SuccessRes(content interface{}) Response {
	return Response{
		Success: true,
		Message: MsgSuccess,
		Content: content,
	}
}