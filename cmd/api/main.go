package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	chMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/lucas-de-lima/gatekeeper/cmd/api/config"
	"github.com/lucas-de-lima/gatekeeper/internal/infra/web/middleware"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Setup structured logging
	setupLogger(cfg)

	slog.Info("Starting Gatekeeper API server",
		"version", cfg.Version,
		"env", cfg.Environment,
		"port", cfg.Server.Port)

	// Initialize router
	r := chi.NewRouter()

	// Global middleware
	r.Use(chMiddleware.Logger)
	r.Use(chMiddleware.Recoverer)
	r.Use(chMiddleware.RequestID)
	r.Use(chMiddleware.RealIP)
	r.Use(middleware.RequestLogger)

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("OK")); err != nil {
			slog.Error("failed to write response", "error", err)
		}
	})

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// TODO: Add API routes here
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			if _, err := w.Write([]byte("Gatekeeper API v1")); err != nil {
				slog.Error("failed to write response", "error", err)
			}
		})
	})

	// Start server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown
	go func() {
		slog.Info("Server starting", "addr", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("Server exited")
}

func setupLogger(cfg *config.Config) {
	var level slog.Level
	switch cfg.Log.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
