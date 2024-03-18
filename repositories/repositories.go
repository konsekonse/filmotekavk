package repositories

import (
	"database/sql"
	"log"
	"vkfilm/models"
)

type DBExecutor interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var db DBExecutor // Используем DBExecutor вместо *sql.DB

// SetDB устанавливает подключение к базе данных для репозиториев
func SetDB(database DBExecutor) {
	db = database
}

// Добавление актёра в базу данных
func AddActor(actor models.Actor) error {
	// Выполнение SQL запроса для добавления актёра
	_, err := db.Exec("INSERT INTO actors (name, gender, date_of_birth) VALUES ($1, $2, $3)",
		actor.Name, actor.Gender, actor.DateOfBirth)
	if err != nil {
		log.Println("Failed to add actor to database:", err)
		return err
	}

	return nil
}

// Обновление информации об актёре в базе данных
func UpdateActor(actorID int, actor models.Actor) error {
	// Выполнение SQL запроса для обновления информации об актёре
	_, err := db.Exec("UPDATE actors SET name = $1, gender = $2, date_of_birth = $3 WHERE id = $4",
		actor.Name, actor.Gender, actor.DateOfBirth, actorID)
	if err != nil {
		log.Println("Failed to update actor in database:", err)
		return err
	}

	return nil
}

// Добавление фильма в базу данных
func AddFilm(film models.Film) error {
	// Выполнение SQL запроса для добавления фильма
	_, err := db.Exec("INSERT INTO films (title, description, release_date, rating) VALUES ($1, $2, $3, $4)",
		film.Title, film.Description, film.ReleaseDate, film.Rating)
	if err != nil {
		log.Println("Failed to add film to database:", err)
		return err
	}

	// Добавление связей актёров с фильмом
	for _, actor := range film.Actors {
		_, err = db.Exec("INSERT INTO film_actors (film_id, actor_id) VALUES (lastval(), $1)", actor.ID)
		if err != nil {
			log.Println("Failed to add actor-film relationship to database:", err)
			return err
		}
	}

	return nil
}

// Обновление информации о фильме в базе данных
func UpdateFilm(filmID int, film models.Film) error {
	// Выполнение SQL запроса для обновления информации о фильме
	_, err := db.Exec("UPDATE films SET title = $1, description = $2, release_date = $3, rating = $4 WHERE id = $5",
		film.Title, film.Description, film.ReleaseDate, film.Rating, filmID)
	if err != nil {
		log.Println("Failed to update film in database:", err)
		return err
	}

	// Удаление связей актёров с фильмом
	_, err = db.Exec("DELETE FROM film_actors WHERE film_id = $1", filmID)
	if err != nil {
		log.Println("Failed to delete actor-film relationships from database:", err)
		return err
	}

	// Добавление связей актёров с фильмом
	for _, actor := range film.Actors {
		_, err = db.Exec("INSERT INTO film_actors (film_id, actor_id) VALUES ($1, $2)", filmID, actor.ID)
		if err != nil {
			log.Println("Failed to add actor-film relationship to database:", err)
			return err
		}
	}

	return nil
}
