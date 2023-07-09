package models

import (
	"github.com/aws/aws-lambda-go/events"
)

type (
	Key string

	RespAPI struct {
		Status     int
		Message    string
		CustomResp *events.APIGatewayProxyResponse
	}
)
