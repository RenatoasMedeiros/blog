package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// we could do this but instead lets build a helper funciton "json"
	//	w.Header().Set("Content-type", "application/json")
	//	w.Write([]byte(`{"status": "ok"}`))
	data := map[string]string{
		"status":  "ok",
		"env":     app.config.env,
		"version": version,
	}
	if err := jsonResponse(w, http.StatusOK, data); err != nil {
		app.internalServerError(w, r, err)
	}
}
