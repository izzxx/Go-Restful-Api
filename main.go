package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/izzxx/Go-Restful-Api/config"
	"github.com/izzxx/Go-Restful-Api/database"
	internal "github.com/izzxx/Go-Restful-Api/middleware"
	"github.com/izzxx/Go-Restful-Api/routes"
	"github.com/allegro/bigcache/v3"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Variable setup
	config.InitConfig()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Minute))
	defer cancel()

	// Manual Database migration
	database.Create(ctx)

	// Postgres
	pgxpool, err := database.OpenDbConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	// BigCache
	memory, err := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Hour))
	if err != nil {
		log.Fatal(err.Error())
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(internal.Limiter())

	timeout := 15 * time.Second

	server := &http.Server{
		Addr:         config.ServerPort,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	// Dependency
	route := &routes.Dependencies{
		Db:     pgxpool,
		Memory: memory,
		App:    e,
	}

	// Route
	route.Health()
	route.User()
	route.Product()

	e.Logger.Fatal(e.StartServer(server))
}
