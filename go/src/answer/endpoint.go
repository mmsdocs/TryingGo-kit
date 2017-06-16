package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeAnswerEndpoint(service AnswerService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		requestObject := request.(answerRequest)

		answer, error_message := service.Answer(requestObject.Message)
		if error_message != nil {
			return answerResponse{answer, error_message.Error()}, nil
		}
		return answerResponse{answer, ""}, nil
	}
}

