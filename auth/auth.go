package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

// Функция для создания токена
func GenerateToken(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		Username: username,
		Role:     role,
	})
	return token.SignedString([]byte("your-secret-key"))
}

// Функция для проверки токена и извлечения информации о пользователе
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
