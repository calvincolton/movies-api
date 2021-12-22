package main

import (
	"encoding/json"
	"errors"
	"log"
	"movies-api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getMovieDetails(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("invalid id parameter"))
		app.errorJSON(w, err)
		return
	}

	movie, err := app.models.DB.Get(id)

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.models.DB.GetAllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) deleteMovie(w http.ResponseWriter, r *http.Request) {}

type MoviePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"releaseDate"`
	Runtime     string `json:"runtime"`
	Rating      string `json:"rating"`
	MPAARating  string `json:"mpaaRating"`
}

func (app *application) insertMovie(w http.ResponseWriter, r *http.Request) {
	var payload MoviePayload

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// log.Println(payload.Title)

	var movie models.Movie

	movie.ID, _ = strconv.Atoi(payload.ID)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.Runtime, _ = strconv.Atoi(payload.Runtime)
	movie.MPAARating = payload.MPAARating
	movie.CreatedAt = time.Now()
	movie.UpdatedAt = time.Now()

	log.Println(movie.Year)

	type jsonResp struct {
		OK bool `json:"ok"`
	}

	ok := jsonResp{
		OK: true,
	}

	err = app.writeJSON(w, http.StatusOK, ok, "response")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

func (app *application) updateMovie(w http.ResponseWriter, r *http.Request) {}

func (app *application) searchMovies(w http.ResponseWriter, r *http.Request) {}

func (app *application) getGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.models.DB.GetAllGenres()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, genres, "genres")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}

// make this return genre name as well as a list of movies by genre id
func (app *application) getMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	genreID, err := strconv.Atoi(params.ByName("genre_id"))
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// genre should return genre details - name, id, etc
	// genre, err := app.models.DB.GetGenre(genreID)
	// if err != nil {
	// 	app.errorJSON(w, err)
	// 	return
	// }

	movies, err := app.models.DB.GetAllMovies(genreID)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// genreDetails := make(map[string]interface)
	// genreDetails := map[string]interface{
	// 	"genre": genre,
	// 	"movies": movies,
	// }

	err = app.writeJSON(w, http.StatusOK, movies, "movies")
	if err != nil {
		app.errorJSON(w, err)
		return
	}
}
