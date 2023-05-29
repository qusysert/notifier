package model

type Message struct {
	Token   string  `json:"token"`
	ChatIds []int64 `json:"chat_ids"`
	Message string  `json:"message"`
}
