package main

import (
	"hello"
	"net/http"
)

const (
	defaultPort = 8080
	defaultRoot = "/"
	isDebug     = true
)

func main() {
	appServerContainer := checkServerStatus()
	appServerContainer.runServer()
}
