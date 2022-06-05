package routes

import "github.com/izzxx/Go-Restful-Api/handler/health"

func (deps *Dependencies) Health() {
	deps.App.GET("/api/v1/ping", health.Health)
}
