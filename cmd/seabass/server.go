package main

import (
	"net/http"
	"strconv"
)

func runServer(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
