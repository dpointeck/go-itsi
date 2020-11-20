package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bitly/go-simplejson"
	"github.com/go-playground/validator"
)

var validate *validator.Validate

func (app *application) TestHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	validate = validator.New()

	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Use the r.PostForm.Get() method to retrieve the relevant data fields
	// from the r.PostForm map.
	title := r.PostForm.Get("title")

	// Initialize a map to hold any validation errors.
	errors := make(map[string]string)

	err = validate.Var(title, "required")
	if err != nil {
		errors["title"] = "This field cannot be blank"
	}

	w.Header().Set("Content-Type", "application/json")

	if len(errors) > 0 {
		json := simplejson.New()
		json.Set("status", "error")
		json.Set("errors", errors)

		payload, err := json.MarshalJSON()
		if err != nil {
			log.Println(err)
		}

		w.Write(payload)
		return
	}

	w.WriteHeader(http.StatusOK)

	io.WriteString(w, fmt.Sprintf(`{"title": "%s"}`, title))

}
