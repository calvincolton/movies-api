package models

import (
	"database/sql"
	"time"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// Movie is the type for movies
type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"releaseDate"`
	Rating      int          `json:"rating"`
	Runtime     int          `json:"runtime"`
	MPAARating  string       `json:"mpaaRating"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	MovieGenre  []MovieGenre `json:"genres"`
}

// Genre is the type for genre
type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genreName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// MovieGenre is the type for movie genre
type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movieId"`
	GenreID   int       `json:"genreId"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
