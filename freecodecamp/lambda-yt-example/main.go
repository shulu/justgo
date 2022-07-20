package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
}

type MyResponse struct {
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
