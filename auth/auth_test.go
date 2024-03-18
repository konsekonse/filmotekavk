package auth_test

import (
	"testing"
	"vkfilm/auth"
)

func TestGenerateToken(t *testing.T) {
	// Генерация токена
	token, err := auth.GenerateToken("user123", "user")
	if err != nil {
		t.Errorf("error generating token: %v", err)
	}

	// Проверка, что токен не пустой
	if token == "" {
		t.Errorf("generated token is empty")
	}
}

func TestParseToken(t *testing.T) {
	// Генерация токена
	token, err := auth.GenerateToken("admin", "admin")
	if err != nil {
		t.Fatalf("error generating token: %v", err)
	}

	// Попытка разбора токена
	claims, err := auth.ParseToken(token)
	if err != nil {
		t.Errorf("error parsing token: %v", err)
	}

	// Проверка, что информация о пользователе извлечена корректно
	if claims.Username != "admin" || claims.Role != "admin" {
		t.Errorf("incorrect claims: got %v, want %v", claims, &auth.Claims{Username: "admin", Role: "admin"})
	}
}
