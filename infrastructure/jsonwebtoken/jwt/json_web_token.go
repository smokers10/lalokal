package jwt

import (
	"encoding/json"
	"fmt"
	"lalokal/infrastructure/configuration"
	"lalokal/infrastructure/jsonwebtoken"

	"github.com/golang-jwt/jwt"
)

type implementation struct {
	applicationConfiguration configuration.Application
}

func JsonWebToken() jsonwebtoken.Contact {
	configuration := configuration.ReadConfiguration()
	return &implementation{applicationConfiguration: configuration.Application}
}

func (i *implementation) Sign(payload map[string]interface{}) (token string, failure error) {
	var covertedPayload jwt.MapClaims

	bytePayload, _ := json.Marshal(payload)
	json.Unmarshal(bytePayload, &covertedPayload)

	tokenResult := jwt.NewWithClaims(jwt.SigningMethodHS256, covertedPayload)
	tokenString, err := tokenResult.SignedString([]byte(i.applicationConfiguration.Secret))
	return tokenString, err
}

func (i *implementation) ParseToken(tokenString string) (payload map[string]interface{}, failure error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(i.applicationConfiguration.Secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return map[string]interface{}{}, err
}
