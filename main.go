package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rzfd/gorm-ners/internal/config"
	"github.com/rzfd/gorm-ners/internal/handlers/http/model"
	"github.com/rzfd/gorm-ners/internal/handlers/http/route"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	config.LoadEnv()

	dbHost := config.GetEnv("DB_HOST")
	dbUser := config.GetEnv("DB_USER")
	dbPassword := config.GetEnv("DB_PASSWORD")
	dbName := config.GetEnv("DB_NAME")
	dbPort := config.GetEnv("DB_PORT")

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&model.User{}, &model.Company{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	e := echo.New()
	route.RegisterRoutes(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}
