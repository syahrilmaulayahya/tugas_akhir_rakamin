package controller

import "fmt"

type BaseResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   []string    `json:"error"`
	Data    interface{} `json:"data"`
}

func (bs *BaseResponse) TokenErrorResp(methodErr string) BaseResponse {
	response := BaseResponse{
		Status:  false,
		Message: fmt.Sprintf("Failed to %s data", methodErr),
		Error:   []string{"Unauthorized"},
		Data:    nil,
	}
	return response
}
