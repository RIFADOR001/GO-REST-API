package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	// The exp key indicates the expiration time of the token
	// In this case we are indicating an expirarion of 2 hours
	// The expiration prevents damage in the case of a stolen token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	// The method used for the token belongs to SigningMethodHMAC
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// This is a way to check if token.Method belongs to the indicated type
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return 0, errors.New("invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}
	// This way we can extract the information from the token
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil
}
