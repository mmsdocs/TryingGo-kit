package main

import (
	"log"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {
	service := answerService{}

	answerHandler := httpTransport.NewServer(
		makeAnswerEndpoint(service),
		decodeAnswerRequest,
		encodeResponse,
	)

	http.Handle("/answerMe", answerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
