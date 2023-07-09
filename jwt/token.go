package jwt

import (
	"errors"
	"strings"

	"github.com/John-Santa/twitter-golang/models"
	"github.com/golang-jwt/jwt/v5"
)

var (
	Email  string
	UserID string
)

func Process(token string, JWTSign string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims models.Claim
	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("token invalido")
	}
	token = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		//TODO: Validar que el token no esté expirado
	}
	if !tkn.Valid {
		return &claims, false, string(""), errors.New("token invalido")
	}
	return &claims, false, string(""), err

}
