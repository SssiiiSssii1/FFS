package main

import (
	"log"

	database "FFS/internal/DB"
	"FFS/internal/config"
	"FFS/internal/feature"
	"FFS/internal/health"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	cfg := config.Load()

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.AutoMigrate(&feature.Feature{}); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	health.RegisterRoutes(router)

	featureRepo := feature.NewRepository(db)
	featureService := feature.NewService(featureRepo)
	featureHandler := feature.NewHandler(featureService)

	feature.RegisterRoutes(router, featureHandler)

	router.Run(":" + cfg.Port)
}
