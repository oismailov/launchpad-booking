package util

type Message struct {
	Message string `json:"message" valid:"required"`
}
