package model

type Response struct {
	Data      interface{} `json:"data"`
	DataCount int         `json:"data_count"`
	Error     string      `json:"error"`
}
