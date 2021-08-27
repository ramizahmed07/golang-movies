package main

import (
	"errors"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIdParam(params httprouter.Params) (int, error) {
	param := params.ByName("id")
	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}
