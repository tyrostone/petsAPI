package main

type Message struct {
	Message string `json: message`
}

func GenerateMessage(text string) Message {
	return Message{Message: text}
}
