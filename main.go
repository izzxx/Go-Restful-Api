package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/izzxx/Go-Restful-Api/config"
	"github.com/izzxx/Go-Restful-Api/database"
	internal "github.com/izzxx/Go-Restful-Api/middleware"
	"github.com/izzxx/Go-Restful-Api/routes"

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
		log.Fatal(err)
	}
	defer pgxpool.Close()

	// BigCache
	memory, err := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Hour))
	if err != nil {
		log.Fatal(err)
	}
	defer memory.Close()

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(internal.Limiter())
	e.Use(middleware.Logger())

	timeout := 15 * time.Second

	server := &http.Server{
		Addr:         ":" + config.ServerPort,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	if server.Addr == ":" {
		server.Addr = ":9000"
	}

	// Dependency
	route := &routes.Dependencies{
		Db:     pgxpool,
		Memory: memory,
		App:    e,
	}

	// Route
	route.User()
	route.Product()
	route.Health()

	e.Logger.Fatal(e.StartServer(server))
}
