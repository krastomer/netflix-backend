package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/krastomer/netflix-backend/database"
	"github.com/krastomer/netflix-backend/handlers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var loggerConfig = middleware.LoggerConfig{
	Format: "${time_rfc3339} | ${remote_ip} | ${status} | ${latency_human} | ${method} ${uri} | ${error}\n",
}

func main() {
	godotenv.Load(".env")
	url := os.Getenv("URL")
	port := os.Getenv("PORT")
	dsn := os.Getenv("DATABASE_URL")

	address := fmt.Sprintf("%s:%s", url, port)

	go database.Initialize(dsn)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(loggerConfig))
	e.Use(middleware.Recover())

	handlers.SetHandlers(e)

	e.Logger.Fatal(e.Start(address))
}
