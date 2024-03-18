package repositories

import (
	"database/sql"
	"errors"
	"testing"
	"vkfilm/models"
)

// MockDB - мок базы данных для тестов
type MockDB struct {
	ExecFunc func(query string, args ...interface{}) (sql.Result, error)
}

// Переопределяем метод Exec для MockDB
func (m *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.ExecFunc != nil {
		return m.ExecFunc(query, args...)
	}
	return nil, errors.New("mocked ExecFunc is not set")
}

func TestAddActor(t *testing.T) {
	// Создаем мок базы данных
	var mockDB DBExecutor = &MockDB{
		ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
			// Проверяем корректность SQL запроса
			expectedQuery := "INSERT INTO actors (name, gender, date_of_birth) VALUES ($1, $2, $3)"
			if query != expectedQuery {
				t.Errorf("unexpected query. got: %s, want: %s", query, expectedQuery)
			}
			// Проверяем переданные аргументы
			expectedArgs := []interface{}{"John Doe", "Male", "1990-01-01"}
			if len(args) != len(expectedArgs) {
				t.Errorf("unexpected arguments. got: %v, want: %v", args, expectedArgs)
			}
			for i, arg := range args {
				if arg != expectedArgs[i] {
					t.Errorf("unexpected argument at index %d. got: %v, want: %v", i, arg, expectedArgs[i])
				}
			}
			// Возвращаем успешный результат
			return nil, nil
		},
	}
	// Устанавливаем мок базы данных
	SetDB(mockDB)

	actor := models.Actor{Name: "John Doe", Gender: "Male", DateOfBirth: "1990-01-01"}
	if err := AddActor(actor); err != nil {
		t.Errorf("AddActor failed with error: %v", err)
	}
}

func TestUpdateActor(t *testing.T) {
	// Создаем мок базы данных
	mockDB := &MockDB{
		ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
			// Проверяем корректность SQL запроса
			expectedQuery := "UPDATE actors SET name = $1, gender = $2, date_of_birth = $3 WHERE id = $4"
			if query != expectedQuery {
				t.Errorf("unexpected query. got: %s, want: %s", query, expectedQuery)
			}
			// Проверяем переданные аргументы
			expectedArgs := []interface{}{"John Doe", "Male", "1990-01-01", 123}
			if len(args) != len(expectedArgs) {
				t.Errorf("unexpected arguments. got: %v, want: %v", args, expectedArgs)
			}
			for i, arg := range args {
				if arg != expectedArgs[i] {
					t.Errorf("unexpected argument at index %d. got: %v, want: %v", i, arg, expectedArgs[i])
				}
			}
			// Возвращаем успешный результат
			return nil, nil
		},
	}
	// Устанавливаем мок базы данных
	SetDB(mockDB)

	actor := models.Actor{Name: "John Doe", Gender: "Male", DateOfBirth: "1990-01-01"}
	if err := UpdateActor(123, actor); err != nil {
		t.Errorf("UpdateActor failed with error: %v", err)
	}
}

func TestAddFilm(t *testing.T) {
	// Создаем мок базы данных
	mockDB := &MockDB{
		ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
			// Проверяем корректность SQL запроса
			expectedQuery := "INSERT INTO films (title, description, release_date, rating) VALUES ($1, $2, $3, $4)"
			if query != expectedQuery {
				t.Errorf("unexpected query. got: %s, want: %s", query, expectedQuery)
			}
			// Возвращаем успешный результат
			return nil, nil
		},
	}
	// Устанавливаем мок базы данных
	SetDB(mockDB)

	film := models.Film{
		Title:       "Film Title",
		Description: "Film Description",
		ReleaseDate: "2023-01-01",
		Rating:      8.5,
		Actors:      []models.Actor{{ID: 1}, {ID: 2}}, // Actors IDs
	}
	if err := AddFilm(film); err != nil {
		t.Errorf("AddFilm failed with error: %v", err)
	}
}

func TestUpdateFilm(t *testing.T) {
	// Создаем мок базы данных
	mockDB := &MockDB{
		ExecFunc: func(query string, args ...interface{}) (sql.Result, error) {
			// Проверяем корректность SQL запроса
			expectedQuery := "UPDATE films SET title = $1, description = $2, release_date = $3, rating = $4 WHERE id = $5"
			if query != expectedQuery {
				t.Errorf("unexpected query. got: %s, want: %s", query, expectedQuery)
			}
			// Возвращаем успешный результат
			return nil, nil
		},
	}
	// Устанавливаем мок базы данных
	SetDB(mockDB)

	film := models.Film{
		Title:       "Film Title",
		Description: "Film Description",
		ReleaseDate: "2023-01-01",
		Rating:      8.5,
		Actors:      []models.Actor{{ID: 1}, {ID: 2}}, // Actors IDs
	}
	if err := UpdateFilm(123, film); err != nil {
		t.Errorf("UpdateFilm failed with error: %v", err)
	}
}
