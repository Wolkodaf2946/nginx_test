package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/wolkodaf/todo/backend/internal/app"
	"github.com/wolkodaf/todo/backend/internal/config"
	"github.com/wolkodaf/todo/backend/pkg/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	application := app.New(log, cfg)

	go func() {
		if err := application.HTTPServer.Run(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("error occured while running http server: %s", err.Error()))
		}
	}()

	log.Info("App Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("App Shutting Down")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := application.HTTPServer.Stop(ctx); err != nil {
		log.Error(fmt.Sprintf("error occured on server shutting down: %s", err.Error()))
	}
	log.Info("Gracefully stopped")
}
