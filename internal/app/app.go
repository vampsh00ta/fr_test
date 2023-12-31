package app

import (
	"context"
	"fmt"
	"fr/internal/repository"
	"fr/internal/service"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"

	"fr/config"
	"fr/internal/transport/http/v1"
	"fr/pkg/client"
	"fr/pkg/httpserver"
	"fr/pkg/logger"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	ctx := context.Background()
	pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		l.Fatal(fmt.Errorf("fr - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	repo := repository.New(pg)

	srvc := service.New(repo)

	handler := gin.New()

	v1.NewRouter(handler, l, srvc)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("fr - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("fr - Run - httpServer.Notify: %w", err))

	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("fr - Run - httpServer.Shutdown: %w", err))
	}

}
