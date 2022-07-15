package models

type Response struct {
	ID      string      `json:"id"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
