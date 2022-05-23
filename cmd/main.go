package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/renoire/shkafchik/src/categories"
	"github.com/renoire/shkafchik/src/httphandlers"
	"github.com/renoire/shkafchik/src/items"
	"github.com/renoire/shkafchik/src/router"
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
	// lf, err := os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0755)
	// if err != nil {
	// 	logrus.Panic("Could not open or create log file: ", err)
	// }

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	cfg := getCfg()

	initLogger(cfg.LogFile)

	h := httphandlers.New(*categories.New(), *items.New())

	srv := http.Server{
		Addr:    cfg.HTTPAddress,
		Handler: router.GetRouter(h),
	}

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
