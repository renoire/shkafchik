package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

type Config struct {
	HTTPAddress string
}

func getCfg() Config {
	httpAddr := os.Getenv("HTTP_ADDRESS")
	if httpAddr == "" {
		httpAddr = "127.0.0.1:8080"
	}

	return Config{
		HTTPAddress: httpAddr,
	}
}

func main() {
	cfg := getCfg()

	srv := http.Server{
		Addr: cfg.HTTPAddress,
	}

	sigC := make(chan os.Signal, 1)
	defer close(sigC)
	go func() {
		<-sigC
		srv.Shutdown(context.TODO()) // nolint:errcheck
	}()

	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("error: http server failed: ", err)
	}

}
