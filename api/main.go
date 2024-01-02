package main

import (
	"context"
	"flag"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/htsuchinga/golang-greeting-api/api/handler/greeting"
	"github.com/htsuchinga/golang-greeting-api/api/internal/server"
	"github.com/htsuchinga/golang-greeting-api/api/internal/validate"
	"github.com/htsuchinga/golang-greeting-api/internal/logger"
)

func main() {
	logger.DefaultModuleName = "api"
	defer func() {
		err := recover()
		if err != nil {
			logger.Error(err)
		}
	}()
	var port uint
	flag.UintVar(&port, "port", 4000, "port")
	flag.Parse()

	// validate
	validator := validate.New()

	srv := server.New(port)

	srv.Handle("/v1/greeting/hello", greeting.NewV1GreetingHandler(validator))

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done()
		logger.Info("Shutting down...")
		if err := srv.Shutdown(); err != nil {
			logger.Warn("server shutdown failed: %s", err)
		}
	}()

	logger.Info("starting server on port %d", port)
	if err := srv.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
