package main

import (
	"database/sql"
	"log"
	"net/http"
	"vkfilm/controllers"
	"vkfilm/repositories"
)

func main() {
	// Создание подключения к базе данных
	db, err := connectToDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Передача подключения к базе данных в репозитории
	repositories.SetDB(db)

	http.HandleFunc("/addactor", controllers.AddActor)
	http.HandleFunc("/updateactor", controllers.UpdateActor)
	http.HandleFunc("/addfilm", controllers.AddFilm)
	http.HandleFunc("/updatefilm", controllers.UpdateFilm)

	// Начало прослушивания запросов
	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectToDatabase() (*sql.DB, error) {

	// Подключение к базе данных
	db, err := sql.Open("postgres", "postgres:admin@tcp(localhost:5042)/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	// Проверка подключения к базе данных
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
