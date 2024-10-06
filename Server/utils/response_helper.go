package utils

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

// レスポンスの構造体を返す
func NewResponse(status int, message string, result interface{}) *response {
	if result == nil {
		// resultがnilの場合、空の配列を設定
		return &response{
			Status:  status,
			Message: message,
			Result:  []interface{}{},
		}
	}
	return &response{
		Status:  status,
		Message: message,
		Result:  result,
	}
}
