package controllers

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// レスポンスの構造体を返す
func newResponse(status int, message string, result interface{}) *response {
    return &response{status, message, result}
}
