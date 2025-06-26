package main

import (
	"appointment-api/internal/api"
	"appointment-api/internal/config"
	"appointment-api/internal/middleware"
	"appointment-api/internal/repository"
	"appointment-api/internal/services"
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration (includes environment loading)
	cfg := config.Load()

	log.Printf("Starting server in %s mode", cfg.App.Environment)
	log.Printf("Database: %s@%s:%s/%s", cfg.Database.User, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	// Database connection
	db, err := sql.Open("postgres", cfg.Database.URL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection successful")

	// Initialize repositories
	repos := repository.NewRepositories(db)

	// Initialize services
	svc := services.NewServices(repos, cfg, db)

	// Start tenant cache
	if err := svc.TenantCache.Start(); err != nil {
		log.Fatal("Failed to start tenant cache:", err)
	}

	// Initialize API handlers
	handlers := api.NewHandlers(svc)

	// Setup router
	if cfg.App.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	// Add CORS middleware
	router.Use(middleware.CORSMiddleware())

	api.SetupRoutes(router, handlers, svc, db, cfg)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", cfg.Server.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Stop tenant cache
	svc.TenantCache.Stop()

	// Shutdown server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server shutdown complete")
}
