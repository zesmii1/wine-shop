package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte("your-secret-key")
	return token.SignedString(secret)
}
