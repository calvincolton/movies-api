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

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		// Adjust status code to 204
		w.WriteHeader(http.StatusNoContent)
	})
	secure := alice.New(app.checkToken)
	// status
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)
	// authentication
	router.HandlerFunc(http.MethodPost, "/v1/signin", app.Signin)
	// for GraphQL
	router.HandlerFunc(http.MethodPost, "/v1/graphql/list", app.moviesGraphQL)

	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.getMovieDetails)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getMovies)
	router.HandlerFunc(http.MethodGet, "/v1/genres", app.getGenres)
	router.HandlerFunc(http.MethodGet, "/v1/genres/:genre_id", app.getMoviesByGenre)

	// secure routes:
	// router.HandlerFunc(http.MethodPost, "/v1/movies", app.insertMovie)
	// router.HandlerFunc(http.MethodPut, "/v1/movies/:id", app.updateMovie)
	// router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovie)
	router.POST("/v1/mvoies", app.wrap(secure.ThenFunc(app.insertMovie)))
	router.PUT("/v1/mvoies", app.wrap(secure.ThenFunc(app.updateMovie)))
	router.DELETE("/v1/mvoies", app.wrap(secure.ThenFunc(app.deleteMovie)))

	return app.enableCORS(router)
}
