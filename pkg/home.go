package home

import (
	"fmt"
	"net/http"
)

// Index is the default action for each URL route in the application
func Index() http.HandlerFunc {
	return (func(rw http.ResponseWriter, rq *http.Request) {
		fmt.Fprintf(rw, "You have reached the homepage of project seabass.")
	})
}
