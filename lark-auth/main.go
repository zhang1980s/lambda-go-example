package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

type Msg struct {
	Challenge string `json:"challenge"`
}

type MsgResponse struct {
	Challenge string `json:"challenge"`
}

func HandleRequest(ctx context.Context, e Msg) (MsgResponse, error) {
	return MsgResponse{
		e.Challenge,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
