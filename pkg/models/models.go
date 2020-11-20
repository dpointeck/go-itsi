package models

import (
	"errors"
	"time"
)

// ErrNoRecord - Error when no record is found
var ErrNoRecord = errors.New("models: no matching record found")

// Book - Type for Book DB Record
type Book struct {
	ID       int
	Title    string    `json:"title"`
	Isbn     string    `json:"isbn"`
	Released time.Time `json:"released"`
	Created  time.Time `json:"created"`
}
