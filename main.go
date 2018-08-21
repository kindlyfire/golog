package main

import (
	"net/http"

	"github.com/kindlyfire/golog/router"
)

func main() {
	// Create web server
	m := router.New()
	router.Register(m)

	// Listen !
	http.ListenAndServe(":6060", m)
}
