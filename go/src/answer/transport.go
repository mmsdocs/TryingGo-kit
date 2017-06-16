package main

import (
	"context"
	"encoding/json"
	"net/http"
)

func decodeAnswerRequest(_ context.Context, httpRequest *http.Request) (interface{}, error) {
	var request answerRequest
	if requestError := json.NewDecoder(httpRequest.Body).Decode(&request); requestError != nil {
		return nil, requestError
	}
	return request, nil
}

func encodeResponse(_ context.Context, responseWriter http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(responseWriter).Encode(response)
}

