package main

import (
	"net/http"

	"github.com/kindlyfire/golog/modules/config"
	"github.com/kindlyfire/golog/router"
)

func main() {
	// I aim to implement reloading, by stopping the web server, then calling start() again
	stoppedChannel := make(chan bool)

	for {
		go start(stoppedChannel)

		<-stoppedChannel
		println("Restarting...")
	}
}

func start(stoppedChannel chan<- bool) {
	config.Load()

	m := router.New()
	srv := &http.Server{Addr: config.Addr, Handler: m}
	router.Register(m)

	srv.ListenAndServe()
}
