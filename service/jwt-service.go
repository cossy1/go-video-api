package service

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken creates a JWT for a given user ID
func GenerateToken(userID uint64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Token expires in 7 days
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtSecretKey)

}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecretKey, nil
	})
}
