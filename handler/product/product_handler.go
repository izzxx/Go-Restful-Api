package product

import (
	"net/http"

	"github.com/izzxx/Go-Restful-Api/service/product"
	"github.com/izzxx/Go-Restful-Api/utility"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductService product.ProductService
}

func (ph *ProductHandler) CreateProduct(c echo.Context) error {
	var storeProduct product.StoreProduct
	err := c.Bind(&storeProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.ErrorResponse(http.StatusBadRequest, err.Error()))
	}

	response, err := ph.ProductService.StoreProduct(c.Request().Context(), storeProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, utility.SuccessResponse("successfully created a new product", response))
}

func (ph *ProductHandler) GetProductById(c echo.Context) error {
	id := c.Param("id")

	responseFromCache, err := ph.ProductService.GetProductIdFromCache(id)
	if responseFromCache != nil || err == nil {
		return c.JSON(http.StatusOK, utility.SuccessResponse("success get product", responseFromCache))
	}

	responseFromDb, err := ph.ProductService.GetProductIdFromDb(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utility.ErrorResponse(http.StatusNotFound, err.Error()))
	}

	return c.JSON(200, utility.SuccessResponse("success get product", responseFromDb))
}

func (ph *ProductHandler) GetAllProducts(c echo.Context) error {
	responseFromCache, err := ph.ProductService.GetAllProductsFromCache()
	if responseFromCache != nil || err == nil {
		return c.JSON(http.StatusOK, utility.SuccessResponse("success get all products", responseFromCache))
	}

	responseFromDb, err := ph.ProductService.GetAllProductsFromDb(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, utility.ErrorResponse(http.StatusNotFound, err.Error()))
	}

	return c.JSON(http.StatusOK, utility.SuccessResponse("success get all products", responseFromDb))
}

func (ph *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	var updateProduct product.UpdateProduct
	err := c.Bind(&updateProduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utility.ErrorResponse(http.StatusBadRequest, err.Error()))
	}
	updateProduct.Id = id

	err = ph.ProductService.UpdateProduct(c.Request().Context(), updateProduct)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utility.ErrorResponse(http.StatusInternalServerError, err.Error()))
	}

	return c.JSON(http.StatusOK, utility.SuccessResponse("success update product", nil))
}

func (ph *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	err := ph.ProductService.DeleteProductById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, utility.ErrorResponse(http.StatusNotFound, err.Error()))
	}

	return c.JSON(http.StatusOK, utility.SuccessResponse("success delete product", nil))
}
