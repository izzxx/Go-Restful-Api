package routes

import (
	handler "github.com/izzxx/Go-Restful-Api/handler/user"
	model "github.com/izzxx/Go-Restful-Api/model/user"
	service "github.com/izzxx/Go-Restful-Api/service/user"
)

func (ud *Dependencies) User() {
	userModels := model.UserRepository{Db: ud.Db}
	userService := service.UserService{UserRepository: userModels}
	userHandler := handler.UserHandler{UserService: userService}

	users := ud.App.Group("/api/v1/users")
	users.POST("/register", userHandler.Register)
	users.POST("/login", userHandler.Login)
}
