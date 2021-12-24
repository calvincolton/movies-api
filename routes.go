package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodPost, "/v1/movies", app.insertMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieDetails)
	router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovie)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getMovies)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getGenres)
	router.HandlerFunc(http.MethodGet, "/v1/genres/:genre_id", app.getMoviesByGenre)

	return app.enableCORS(router)
}
