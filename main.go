package main

import (
	"net/http"

	"github.com/kindlyfire/golog/modules/config"
	"github.com/kindlyfire/golog/modules/db"
	"github.com/kindlyfire/golog/modules/theme"
	"github.com/kindlyfire/golog/router"
)

func main() {
	// I aim to implement reloading, by stopping the web server, then calling start() again
	stoppedChannel := make(chan bool)

	for {
		go start(stoppedChannel)

		// Wait for start() to signal us it has shut down
		<-stoppedChannel

		println("Restarting...")
	}
}

func start(stoppedChannel chan<- bool) {
	// Load website config.ini
	config.Load()

	// Connect to the database
	db.Connect()

	// Load theme
	theme.Load()

	// Create router & server
	m := router.New()
	srv := &http.Server{Addr: config.Addr, Handler: m}

	// Register all routes
	router.Register(m)

	// Listen !
	srv.ListenAndServe()
}
