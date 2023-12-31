package handlers

import (
	"context"
	"fmt"

	"github.com/John-Santa/twitter-golang/models"
	"github.com/aws/aws-lambda-go/events"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RespAPI {
	fmt.Println("> Inicio de procesamiento de la petición a", ctx.Value(models.Key("path")).(string), ">", ctx.Value(models.Key("method")).(string))

	var r models.RespAPI
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}
	r.Message = "Método no soportado"
	return r
}
