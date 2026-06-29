package main

import (
	"log"
	"os"

	"smart-aquaculture-backend/delivery/http"
	"smart-aquaculture-backend/infrastructure"
	"smart-aquaculture-backend/repository"
	"smart-aquaculture-backend/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file if exists
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	// 1. Init Database (PostgreSQL / Supabase)
	db := infrastructure.ConnectDB()

	// 2. Init Repository
	monitoringRepo := repository.NewMonitoringRepository(db)
	tokenRepo := repository.NewTokenRepository(db)

	// Init Firebase
	infrastructure.InitFirebase()

	// 3. Init Usecase
	monitoringUsecase := usecase.NewMonitoringUsecase(monitoringRepo, tokenRepo)

	// 4. Init MQTT Background Worker
	mqttClient := infrastructure.InitMQTT(monitoringUsecase.SaveTelemetryData)
	defer mqttClient.Disconnect(250)

	// 5. Setup Fiber HTTP Server
	app := fiber.New()
	app.Use(cors.New()) // Enable CORS for Frontend

	api := app.Group("/api/v1")
	
	monitoringGroup := api.Group("/monitoring")
	http.NewMonitoringHandler(monitoringGroup, monitoringUsecase)

	controlGroup := api.Group("/control")
	http.NewControlHandler(controlGroup, mqttClient)

	// FCM Token Registration
	notiHandler := http.NewNotificationHandler(tokenRepo)
	api.Post("/fcm/token", notiHandler.SaveToken)

	// 6. Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Smart Aquaculture Backend is running on port %s...", port)
	log.Fatal(app.Listen(":" + port))
}
