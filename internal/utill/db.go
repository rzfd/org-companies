package utill

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/rzfd/gorm-ners/internal/config"
	"github.com/rzfd/gorm-ners/internal/handlers/http/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() *gorm.DB {
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

	if err := db.AutoMigrate(&entities.User{}, &entities.Company{}, &entities.Regis{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	return db
}
