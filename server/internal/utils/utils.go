package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/personal-project/zentio/internal/config"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

func InitJwtSecret() {
	secretKey, err := config.GetEnv("JWT_SECRET")
	if err != nil {
		log.Fatal("Jwt secret is not defined")
	}
	jwtSecret = []byte(secretKey)
}

func GenerateToken(userId uint, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"iat":      expirationTime,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}
