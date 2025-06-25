package main

import (
	"appointment-api/internal/api"
	"appointment-api/internal/config"
	"appointment-api/internal/middleware"
	"appointment-api/internal/repository"
	"appointment-api/internal/services"
	"database/sql"
	"log"

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

	// Start server
	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Fatal(router.Run(":" + cfg.Server.Port))
}
