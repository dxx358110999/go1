package gin_middleware

const AppCodeSuccess int = 200

type ResponseBody struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
