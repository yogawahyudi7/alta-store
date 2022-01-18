package common

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func SuccessResponse(data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Code:    200,
		Message: "success",
		Data:    data,
	}
}

func ErrorResponse(code int, message string) ResponseError {
	return ResponseError{
		Code:    code,
		Message: message,
	}
}
