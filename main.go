package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
)

type Email struct {
	RecepientEmail string `validate:"required,email"`
	Content        string `validate:"required"`
}

var Validate = validator.New()

func main() {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/health", func(rw http.ResponseWriter, r *http.Request) {

		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte("I am a big, strong, healthy boy."))
	})

	r.Post("/send-mail", func(rw http.ResponseWriter, req *http.Request) {

		var email Email

		decoder := json.NewDecoder(req.Body)

		var err = decoder.Decode(&email)
		if err != nil {
			http.Error(rw, "Bad Request : "+err.Error(), http.StatusBadRequest)
			return
		}

		err = Validate.Struct(email)

		if err != nil {
			http.Error(rw, "Bad Request : "+err.Error(), http.StatusBadRequest)
			return
		}

	})

	http.ListenAndServe(":3001", r)

}
