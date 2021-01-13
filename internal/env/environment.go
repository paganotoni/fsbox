package env

import (
	"os"
)

// These is the place we put constants that will be used across the app
// to start these 3 constants come very handy to know which environment
// are you in.
const (
	Development = "development"
	Tests       = "tests"
	Production  = "production"
)

// Current returns the name of the environment stored in GO_ENV
// and defaults to development if empty.
func Current() string {
	e := os.Getenv("GO_ENV")
	if e == "" {
		return Development
	}

	return e
}
