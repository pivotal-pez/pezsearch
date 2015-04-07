package pezsearch

const (
	successStatus = "success"
	errorStatus   = "error"
)

//ResponseMessage structures output into a standard format.
type ResponseMessage struct {
	Status  string                 `json:"status"`
	Data    interface{}            `json:"data,omitempty"`
	Message string                 `json:"message,omitempty"`
	Meta    map[string]interface{} `json:"_meta,omitempty"`
	Links   map[string]interface{} `json:"_links,omitempty"`
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
