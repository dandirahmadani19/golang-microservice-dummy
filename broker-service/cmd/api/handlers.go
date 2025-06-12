package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the Broker api",
	}

	err := app.writeJSON(w, http.StatusAccepted, payload)
	if err != nil {
		panic(err)
	}
}
