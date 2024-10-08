package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rzfd/gorm-ners/internal/config"
	"github.com/rzfd/gorm-ners/internal/handlers/http/controller"
	"github.com/rzfd/gorm-ners/internal/handlers/http/repositories"
	"github.com/rzfd/gorm-ners/internal/handlers/http/route"
	"github.com/rzfd/gorm-ners/internal/handlers/http/security"
	"github.com/rzfd/gorm-ners/internal/handlers/http/services"
	"github.com/rzfd/gorm-ners/internal/utill"
)

func main() {
	config.LoadEnv()
	dbConn := utill.ConnectDB()

	userRepo := repositories.NewUserRepository(dbConn)
	authService := services.NewAuthService(userRepo)

	e := echo.New()
	jwtSecret := os.Getenv("JWT_TOKEN")

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	e.POST("/register", controller.Register(authService))
	e.POST("/login", controller.Login(authService))

	protect := e.Group("")
	protect.Use(security.JWTMiddleware(jwtSecret))

	route.RegisterRoutes(protect, dbConn)
	e.Logger.Fatal(e.Start(":8080"))
}
