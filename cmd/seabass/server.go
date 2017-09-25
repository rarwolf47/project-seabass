package main

import (
	"net/http"
	"strconv"
)

func RunServer(port int) {
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
