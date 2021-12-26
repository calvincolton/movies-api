package main

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) wrap(next http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := context.WithValue(r.Context(), "params", ps)
		// w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	secure := alice.New(app.checkToken)

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)

	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieDetails)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getMovies)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getGenres)
	router.HandlerFunc(http.MethodGet, "/v1/genres/:genre_id", app.getMoviesByGenre)

	// secure routes:
	router.POST("/v1/mvoies", app.wrap(secure.ThenFunc(app.insertMovie)))
	router.PUT("/v1/mvoies", app.wrap(secure.ThenFunc(app.updateMovie)))
	router.DELETE("/v1/mvoies", app.wrap(secure.ThenFunc(app.deleteMovie)))
	// router.HandlerFunc(http.MethodPost, "/v1/movies", app.insertMovie)
	// router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovie)
	// router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovie)

	return app.enableCORS(router)
}
