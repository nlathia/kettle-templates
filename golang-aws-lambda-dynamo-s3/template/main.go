package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	// Errors are ignored here for now
	dao.InitSession()
	storage.InitSession()
}

func LambdaHandler(req domain.RequestEvent) (domain.ResponseEvent, error) {
	if err := domain.ValidateRequest(req); err != nil {
		return domain.ResponseEvent{}, err
	}
	if err := dao.InitSession(); err != nil {
		return domain.ResponseEvent{}, err
	}
	if err := storage.InitSession(); err != nil {
		return domain.ResponseEvent{}, err
	}

	return domain.ResponseEvent{
		Message: fmt.Sprintf("Hello, %s", req.UserName),
	}, nil
}

func main() {
	lambda.Start(LambdaHandler)
}
