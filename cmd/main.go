package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"v002_onelab/configs"
	"v002_onelab/internal/repository"
	"v002_onelab/internal/service"
	rest "v002_onelab/internal/transport/http"

	_ "v002_onelab/docs"
)

// @title Transaction
// @version 1.0
// @description REST API Library

// @contact.name Serikov Dias
// @contact.email serikov.2002.12@gmail.com

// @host localhost:8090
// @BasePath /api

// @securityDefinitions.apiKey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	log.Fatal(run())
}

func run() error {
	config, err := configs.New()
	if err != nil {
		return err
	}

	repo, err := repository.New(config)
	if err != nil {
		return err
	}
	service := service.New(repo)
	handler := rest.New(service, config)
	srv := handler.InitRouter()

	go func() {
		if err := srv.Start(fmt.Sprintf(":%s", config.PORT)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	gracefulShutdown(srv)
	return nil
}

func gracefulShutdown(srv *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Println("Server stopped gracefully")
}
