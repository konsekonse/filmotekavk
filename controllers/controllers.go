package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"vkfilm/auth"
	"vkfilm/models"
	"vkfilm/repositories"
)

// Контроллер для добавления актёра
func AddActor(w http.ResponseWriter, r *http.Request) {
	// Проверка авторизации администратора
	claims, err := auth.ParseToken(r.Header.Get("Authorization"))
	if err != nil || claims.Role != "admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Парсинг данных об актёре из запроса
	var actor models.Actor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Добавление актёра в базу данных
	err = repositories.AddActor(actor)
	if err != nil {
		http.Error(w, "Failed to add actor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Контроллер для изменения информации об актёре
func UpdateActor(w http.ResponseWriter, r *http.Request) {
	// Проверка авторизации администратора
	claims, err := auth.ParseToken(r.Header.Get("Authorization"))
	if err != nil || claims.Role != "admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Извлечение ID актёра из запроса
	actorID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid actor ID", http.StatusBadRequest)
		return
	}

	// Парсинг данных об актёре из запроса
	var actor models.Actor
	if err := json.NewDecoder(r.Body).Decode(&actor); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Обновление информации об актёре
	err = repositories.UpdateActor(actorID, actor)
	if err != nil {
		http.Error(w, "Failed to update actor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Контроллер для добавления фильма
func AddFilm(w http.ResponseWriter, r *http.Request) {
	// Проверка авторизации администратора
	claims, err := auth.ParseToken(r.Header.Get("Authorization"))
	if err != nil || claims.Role != "admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Парсинг данных о фильме из запроса
	var film models.Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Добавление фильма в базу данных
	err = repositories.AddFilm(film)
	if err != nil {
		http.Error(w, "Failed to add film", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Контроллер для изменения информации о фильме
func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	// Проверка авторизации администратора
	claims, err := auth.ParseToken(r.Header.Get("Authorization"))
	if err != nil || claims.Role != "admin" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Извлечение ID фильма из запроса
	filmID, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid film ID", http.StatusBadRequest)
		return
	}

	// Парсинг данных о фильме из запроса
	var film models.Film
	if err := json.NewDecoder(r.Body).Decode(&film); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Обновление информации о фильме
	err = repositories.UpdateFilm(filmID, film)
	if err != nil {
		http.Error(w, "Failed to update film", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
