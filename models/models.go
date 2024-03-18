package models

type Actor struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	DateOfBirth string `json:"date_of_birth"`
}

type Film struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"release_date"`
	Rating      float64 `json:"rating"`
	Actors      []Actor `json:"actors"`
}
