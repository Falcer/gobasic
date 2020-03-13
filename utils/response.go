package utils

// Response base that we gonna send to client
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var response Response

// SuccessResponse is we succesed process all stuffs that client wanna to use
func SuccessResponse(message string, data interface{}) *Response {
	response.Status = "success"
	response.Message = message
	response.Data = data
	return &response
}

// ErrorResponse is we succesed process all stuffs that client wanna to use
func ErrorResponse(message string, data interface{}) *Response {
	response.Status = "error"
	response.Message = message
	response.Data = data
	return &response
}
