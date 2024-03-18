package controllers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"vkfilm/auth"
	"vkfilm/controllers"
	"vkfilm/models"
	"vkfilm/repositories"
)

// Мок для реализации интерфейса DBExecutor
type mockDBExecutor struct{}

func (m *mockDBExecutor) Exec(query string, args ...interface{}) (sql.Result, error) {
	// Возвращаем успешный результат
	return nil, nil
}

func TestAddActor(t *testing.T) {
	// Установка мок-объекта в качестве базы данных
	repositories.SetDB(&mockDBExecutor{})

	// Создание запроса
	reqBody := models.Actor{Name: "John Doe", Gender: "Male", DateOfBirth: "1990-01-01"}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/actors", bytes.NewReader(reqBodyBytes))

	// Генерация токена
	token, err := auth.GenerateToken("admin", "admin")
	if err != nil {
		t.Fatal("Failed to generate token:", err)
	}

	// Заголовок авторизации
	req.Header.Set("Authorization", token)

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов контроллера
	controllers.AddActor(rr, req)

	// Проверка кода состояния
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestAddFilm(t *testing.T) {
	// Установка мок-объекта в качестве базы данных
	repositories.SetDB(&mockDBExecutor{})

	// Создание запроса
	reqBody := models.Film{Title: "Sample Film", Description: "Sample Description", ReleaseDate: "2023-01-01", Rating: 8.5}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/films", bytes.NewReader(reqBodyBytes))

	// Генерация токена
	token, err := auth.GenerateToken("admin", "admin")
	if err != nil {
		t.Fatal("Failed to generate token:", err)
	}

	// Заголовок авторизации
	req.Header.Set("Authorization", token)

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов контроллера
	controllers.AddFilm(rr, req)

	// Проверка кода состояния
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
}

func TestUpdateFilm(t *testing.T) {
	// Установка мок-объекта в качестве базы данных
	repositories.SetDB(&mockDBExecutor{})

	// Создание запроса
	reqBody := models.Film{Title: "Updated Film", Description: "Updated Description", ReleaseDate: "2023-01-01", Rating: 9.0}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("PUT", "/films?id=123", bytes.NewReader(reqBodyBytes))

	// Генерация токена
	token, err := auth.GenerateToken("admin", "admin")
	if err != nil {
		t.Fatal("Failed to generate token:", err)
	}

	// Заголовок авторизации
	req.Header.Set("Authorization", token)

	// Создание ResponseRecorder для записи ответа
	rr := httptest.NewRecorder()

	// Вызов контроллера
	controllers.UpdateFilm(rr, req)

	// Проверка кода состояния
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
