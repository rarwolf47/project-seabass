package main

import (
	"net/http"
	"strconv"
)

type (
	applicationServerContainer interface {
		runServer()
	}

	offlineServerContainer struct {
		port      int
		rootRoute string
	}

	onlineServerContainer struct {
		port      int
		rootRoute string
	}
)

func checkServerStatus() applicationServerContainer {
	if isDebug {
		return newOfflineServerContainer()
	}

	return newOnlineServerContainer()
}

func newOfflineServerContainer() offlineServerContainer {
	var server offlineServerContainer
	server.port = defaultPort
	server.rootRoute = defaultRoot
	return server
}

func newOnlineServerContainer() onlineServerContainer {
	var server onlineServerContainer
	server.port = defaultPort
	server.rootRoute = defaultRoot
	return server
}

func (server offlineServerContainer) runServer() {
	http.ListenAndServe(":"+strconv.Itoa(server.port), nil)
}

func (server onlineServerContainer) runServer() {
	http.ListenAndServe(":"+strconv.Itoa(server.port), nil)
}
