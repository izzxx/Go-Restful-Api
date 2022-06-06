package product

import (
	"context"
	"errors"

	"github.com/pquerna/ffjson/ffjson"
)

// From database
func (ps *ProductService) GetProductIdFromDb(ctx context.Context, id string) (*ProductResponse, error) {
	if id == "" {
		return nil, errors.New("id not found")
	}

	product, err := ps.ProductModel.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	response := ProductResponse{
		Id:         product.Id,
		Name:       product.Name,
		Price:      product.Price,
		Quantity:   product.Quantity,
		Created_At: product.Created_At,
	}

	return &response, nil
}

func (ps *ProductService) GetAllProductsFromDb(ctx context.Context) ([]ProductResponse, error) {
	products, err := ps.ProductModel.FindAllProduct(ctx)
	if err != nil {
		return nil, err
	}

	// nil slice
	// if all products are empty, it will be return nil
	var productsResponse []ProductResponse

	// empty slice
	// var productsResponse = []ProductResponse{}
	// productsResponse := []ProductsResponse{}

	for _, product := range products {
		productsResponse = append(productsResponse, ProductResponse(product))
	}

	if productsResponse == nil {
		return nil, errors.New("products not found")
	}

	if err = ps.SetCacheProducts(productsResponse); err != nil {
		return nil, err
	}

	return productsResponse, nil
}

// From memory
func (ps *ProductService) GetProductIdFromCache(id string) (*ProductResponse, error) {
	products, err := ps.GetAllProductsFromCache()
	if err != nil {
		return nil, err
	}

	for _, product := range products {
		if product.Id == id {
			return &product, nil
		}
	}

	return nil, errors.New("product with id " + id + " not found")
}

func (ps *ProductService) GetAllProductsFromCache() ([]ProductResponse, error) {
	cache, err := ps.Memory.Get("products")
	if err != nil {
		return nil, err
	}

	// nil slice
	var products []ProductResponse
	if err = ffjson.Unmarshal(cache, &products); err != nil {
		return nil, err
	}

	return products, nil
}
