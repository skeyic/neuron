package model

type PostMan interface {
	GetID() string
	Send(body *AlertBody) error
}

type AlertBody struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
