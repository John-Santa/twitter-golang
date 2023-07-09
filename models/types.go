package models

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Key string

	RespAPI struct {
		Status     int
		Message    string
		CustomResp *events.APIGatewayProxyResponse
	}

	Claim struct {
		Email string             `json:"email"`
		ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
		jwt.RegisteredClaims
	}
)
