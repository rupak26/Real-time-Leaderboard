package domain

type ApiResponse struct {
	Success    bool        `json:"success"`
	Message    interface{} `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	DataList   interface{} `json:"dataList,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
}

