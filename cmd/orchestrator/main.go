package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/wisaitas/grpc-poc/internal/orchestrator/initial"
)

func main() {
	app := initial.New()
	defer app.Stop()

	go func() {
		app.Start()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Stop()
}
