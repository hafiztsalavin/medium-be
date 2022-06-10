package utils

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseLoginSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   interface{} `json:"token"`
}

func SuccessResponse(data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Code:    200,
		Message: "Successful Operation",
		Data:    data,
	}
}

func SuccessLoginResponse(data, token interface{}) ResponseLoginSuccess {
	return ResponseLoginSuccess{
		Code:    200,
		Message: "Successful Operation",
		Data:    data,
		Token:   token,
	}
}

func ErrorResponse(code int, message string) DefaultResponse {
	return DefaultResponse{
		Code:    code,
		Message: message,
	}
}

//NewSuccessOperationResponse default success operation response
func NewSuccessOperationResponse() DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation",
	}
}

//NewNotFoundResponse default not found error response
func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		404,
		"Not Found",
	}
}

//NewBadRequestResponse default bad request error response
func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		400,
		"Bad Request",
	}
}

//NewStatusNotAccepted default not
func NewStatusNotAcceptable() DefaultResponse {
	return DefaultResponse{
		406,
		"Not Accepted",
	}
}

//NewUnauthorize default Unauthorize response
func NewUnauthorizeResponse() DefaultResponse {
	return DefaultResponse{
		401,
		"Unauthorize",
	}
}
