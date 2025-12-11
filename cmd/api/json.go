package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// special function in go, runs this whenever the program initializes
func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled()) //recomended so use it
}

// http handler (everything on the api folder is http handlers so coupling it with http.ResponseWriter is not a problem)
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 // 1 MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	//disallow unknown fields
	decoder.DisallowUnknownFields() // this will thow an error if a unknown field appear
	return decoder.Decode(data)
}

// Why do i build a function for this?
// the api responses should be consistent, any client of this api should have consistent error messages!
func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	//pointer to a envelope that as the fields that we want to return
	return writeJSON(w, status, &envelope{Error: message})
}

func jsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	return writeJSON(w, status, &envelope{Data: data})
}
