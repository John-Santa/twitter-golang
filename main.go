package twitter_golang

import (
	"context"
	"os"
	"strings"

	"github.com/John-Santa/twitter-golang/awsgo"
	"github.com/John-Santa/twitter-golang/models"
	"github.com/John-Santa/twitter-golang/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	//BOOKMARK: Llamada o puerta de entrada de la lambda
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var response *events.APIGatewayProxyResponse
	awsgo.AwsInit()

	if !ValidateParams() {
		response = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error: Se deben configurar las variables de entorno SecretName, BucketName y UrlPrefix",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return response, nil
	}
	secret, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		response = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura del secret: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return response, nil
	}

	path := strings.Replace(request.PathParameters["twitter-go-backend"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), secret.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), secret.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), secret.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), secret.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwt_sign"), secret.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucket_name"), os.Getenv("BucketName"))
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)

	return nil, nil
}

func ValidateParams() bool {
	_, exist := os.LookupEnv("SecretName")
	if !exist {
		return exist
	}
	_, exist = os.LookupEnv("bucketName")
	if !exist {
		return exist
	}
	_, exist = os.LookupEnv("UrlPrefix")
	if !exist {
		return exist
	}
	return exist
}
