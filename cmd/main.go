package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

type Config struct {
	HTTPAddress string
	LogFile     string
}

func getCfg() Config {
	httpAddr := os.Getenv("HTTP_ADDRESS")
	if httpAddr == "" {
		httpAddr = "127.0.0.1:8080"
	}

	logFile := os.Getenv("LOG_FILE")
	if logFile == "" {
		logFile = "log.txt"
	}

	return Config{
		HTTPAddress: httpAddr,
		LogFile:     logFile,
	}
}

func initLogger(logFile string) {
	lf, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logrus.Panic("Could not open or create log file: ", err)
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(lf)
}

func main() {
	cfg := getCfg()

	srv := http.Server{
		Addr: cfg.HTTPAddress,
	}

	initLogger(cfg.LogFile)

	sigC := make(chan os.Signal, 1)
	defer close(sigC)
	go func() {
		<-sigC
		srv.Shutdown(context.TODO()) // nolint:errcheck
	}()

	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.Error("error: http server failed: ", err)
	}

}
