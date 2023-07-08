package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/John-Santa/twitter-golang/awsgo"
	"github.com/John-Santa/twitter-golang/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secret models.Secret
	fmt.Println("> Solicitando el secreto: ", secretName)
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	secrets, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println("> Error al solicitar el secreto: ", err.Error())
		return secret, err
	}
	json.Unmarshal([]byte(*secrets.SecretString), &secret)
	fmt.Println("> Lectura del secreto exitosa ", secretName)
	return secret, nil
}
