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

	ud.App.POST("/api/v1/users/register", userHandler.Register)
	ud.App.POST("/api/v1/users/login", userHandler.Login)
}
