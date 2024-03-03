package models

type Error struct {
	Error error `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
