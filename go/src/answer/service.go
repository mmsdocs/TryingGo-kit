package main

import (
	"errors"
	"strings"
)

type AnswerService interface {
	Answer(string) (string, error)
}

type answerService struct{}

func (answerService) Answer(message string) (string, error) {
	message = strings.Replace(message, "0", "o", -1)
	message = strings.Replace(message, "3", "e", -1)
	message = strings.Replace(message, "1", "l", -1)

	if strings.ToLower(message) != "hello" {
		return "I don't understand you!", ErrorUnderstand
	}
	return "Oh, Hello! How are you?", nil
}

var ErrorUnderstand = errors.New("Try: Hello")

