package models

type StandartError struct {
	Error error `json:"error"`
}

type ForbiddenError struct {
	Message string `json:"message"`
}
