package main

type Error struct {
	Message string `json: message`
}

func GenerateError(text string) Error {
	return Error{Message: text}
}
