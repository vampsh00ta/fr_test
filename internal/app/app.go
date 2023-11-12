package app

import (
	"context"
	"fmt"
	"fr/cmd/fr/docs"
	"fr/internal/repository"
	"fr/internal/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
	"os/signal"
	"syscall"

	"fr/config"
	"fr/internal/transport/http/v1"
	"fr/pkg/client"
	"fr/pkg/httpserver"
	"fr/pkg/logger"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	ctx := context.Background()
	// Repository
	pg, err := client.NewPostgresClient(ctx, 5, cfg.PG)
	if err != nil {
		l.Fatal(fmt.Errorf("fr - Run - postgres.New: %w", err))
	}
	defer pg.Close()
	repo := repository.New(pg)

	// Service
	srvc := service.New(repo)
	err = srvc.SendNewsletter(ctx, 4)

	fmt.Println(err)
	// HTTP Server
	handler := gin.New()
	docs.SwaggerInfo.BasePath = "/v1/"
	handler.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1.NewRouter(handler, l, srvc)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("fr - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("fr - Run - httpServer.Notify: %w", err))

	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("fr - Run - httpServer.Shutdown: %w", err))
	}

}
