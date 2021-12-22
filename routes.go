package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieDetails)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getMovies)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getGenres)
	router.HandlerFunc(http.MethodGet, "/v1/genres/:genre_id", app.getMoviesByGenre)
	router.HandlerFunc(http.MethodPost, "/v1/movies/new", app.insertMovie)

	return app.enableCORS(router)
}
