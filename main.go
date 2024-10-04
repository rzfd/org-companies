package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rzfd/gorm-ners/internal/config"
	"github.com/rzfd/gorm-ners/internal/handlers/http/controller"
	"github.com/rzfd/gorm-ners/internal/handlers/http/middleware"
	"github.com/rzfd/gorm-ners/internal/handlers/http/route"
	"github.com/rzfd/gorm-ners/internal/utill"
)

func main() {
	config.LoadEnv()
	dbConn := utill.ConnectDB()
	e := echo.New()
	jwtSecret := os.Getenv("JWT_SECRET")
	e.POST("/register", controller.Regis(dbConn, jwtSecret))
	e.POST("/login", controller.Login(dbConn, jwtSecret))
	protect := e.Group("")
	protect.Use(middleware.JWTMiddleware(jwtSecret))
	route.RegisterRoutes(e, dbConn)
	e.Logger.Fatal(e.Start(":8080"))
}
