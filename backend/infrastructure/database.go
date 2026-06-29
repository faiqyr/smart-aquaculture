package infrastructure

import (
	"fmt"
	"log"
	"os"

	"smart-aquaculture-backend/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB establishes connection to Supabase PostgreSQL and runs AutoMigrate
func ConnectDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
    
	if host == "" {
		log.Fatal("DB_HOST environment variable is required")
	}

	// 🛠️ PERBAIKAN: Tambahkan default_query_exec_mode=simple_protocol di ujung string DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require default_query_exec_mode=simple_protocol",
		host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false, // Mematikan cache statement di level GORM
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Auto Migrate the models
	err = db.AutoMigrate(&domain.PoolMonitoring{}, &domain.DeviceControl{}, &domain.DeviceToken{})
	if err != nil {
		log.Fatalf("Failed to auto migrate database: %v", err)
	}

	log.Println("Database connection successfully opened and migrated")
	return db
} 