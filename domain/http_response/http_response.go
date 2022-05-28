package http_response

type Response struct {
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success"`
	Status  int         `json:"status,omitempty"`
	Token   string      `json:"token,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
