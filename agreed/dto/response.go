package dto

type ResponseBody struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
