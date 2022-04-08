package model

type BaseResponse struct {
	Status int    `json:"status"`
	Info   string `json:"info"`
}
