package models

import "time"

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"releaseDate"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaaRating"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	MovieGenre  []MovieGenre `json:"-"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genreName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type MovieGenre struct {
	ID        int       `json:"id"`
	MovieID   int       `json:"movieId"`
	GenreId   int       `json:"genreId"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
