package controllers

const (
	successStatus = "success"
	errorStatus   = "error"
	failStatus    = "fail"
)

//ResponseMessage structures output into a standard format.
type ResponseMessage struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func successMessage(data interface{}) (rsp ResponseMessage) {
	rsp = ResponseMessage{
		Status: successStatus,
		Data:   data,
	}
	return
}

func errorMessage(message string) (rsp ResponseMessage) {
	rsp = ResponseMessage{
		Status:  errorStatus,
		Message: message,
	}
	return
}

func failureMessage(data interface{}) (rsp ResponseMessage) {
	rsp = ResponseMessage{
		Status: failStatus,
		Data:   data,
	}
	return
}
