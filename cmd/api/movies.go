package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/ramizahmed07/golang-movies/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := app.readIdParam(params)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        int64(id),
		CreatedAt: time.Now(),
		Title:     "Interstellar",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "mystery"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
