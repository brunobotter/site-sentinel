package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	deliveryhttp "github.com/seuuser/go-site-monitor/internal/delivery/http"
	"github.com/seuuser/go-site-monitor/internal/domain"
	"github.com/seuuser/go-site-monitor/internal/repository"
	"github.com/seuuser/go-site-monitor/internal/service"
	"github.com/seuuser/go-site-monitor/pkg/logger"
)

func main() {
	log := logger.New()

	siteRepo := repository.NewMemorySiteRepository([]domain.Site{
		{ID: "1", URL: "https://google.com", IntervalSeconds: 60},
		{ID: "2", URL: "https://github.com", IntervalSeconds: 60},
	})
	_ = repository.NewMemoryCheckRepository()

	siteService := service.NewSiteService(siteRepo)
	router := deliveryhttp.NewRouter(siteService)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("HTTP server running on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
		return
	}

	log.Println("server stopped")
}
