package main

type answerRequest struct {
	Message string `json:"message"`
}

type answerResponse struct {
	AnswerData   string `json:"answer"`
	ErrorMessage string `json:"error_message,omitempty"`
}

