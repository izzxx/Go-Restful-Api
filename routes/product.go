package routes

import (
	handler "github.com/izzxx/Go-Restful-Api/handler/product"
	"github.com/izzxx/Go-Restful-Api/middleware"
	model "github.com/izzxx/Go-Restful-Api/model/product"
	service "github.com/izzxx/Go-Restful-Api/service/product"
)

func (deps *Dependencies) Product() {
	productModel := model.ProductModel{Db: deps.Db}
	productService := service.ProductService{ProductModel: productModel, Memory: deps.Memory}
	productHandler := handler.ProductHandler{ProductService: productService}

	// Must be registered as a user
	productForUser := deps.App.Group("/api/v1", middleware.AuthMiddleware)
	productForUser.GET("/products", productHandler.GetAllProducts)
	productForUser.GET("/products/:id", productHandler.GetProductById)

	// Must be registered as admin
	productAccessForAdmin := productForUser.Group("/admin", middleware.IsAdmin)
	productAccessForAdmin.POST("/products", productHandler.CreateProduct)
	productAccessForAdmin.PUT("/products/:id", productHandler.UpdateProduct)
	productAccessForAdmin.DELETE("/products/:id", productHandler.DeleteProduct)
}
