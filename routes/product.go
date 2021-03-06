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
	productAccessForUser := deps.App.Group("/api/v1/products", middleware.AuthMiddleware)
	productAccessForUser.GET("", productHandler.GetAllProducts)
	productAccessForUser.GET("/:id", productHandler.GetProductById)

	// Must be registered as admin
	productAccessForAdmin := productAccessForUser.Group("/admin", middleware.IsAdmin)
	productAccessForAdmin.POST("/create", productHandler.CreateProduct)
	productAccessForAdmin.PUT("/update/:id", productHandler.UpdateProduct)
	productAccessForAdmin.DELETE("/delete/:id", productHandler.DeleteProduct)
}
