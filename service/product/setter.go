package product

import (
	"context"
	"errors"

	"github.com/izzxx/Go-Restful-Api/config"
	"github.com/izzxx/Go-Restful-Api/model/product"
	"github.com/izzxx/Go-Restful-Api/utility"
	"github.com/pquerna/ffjson/ffjson"
)

// When data is changed or delete, you should also clear the cache

func (ps *ProductService) StoreProduct(ctx context.Context, prod StoreProduct) (*ProductResponse, error) {
	switch {
	case len(prod.Name) < 2:
		return nil, errors.New("product name letters cannot be less than 2")
	case prod.Price < 0:
		return nil, config.ErrorProductPrice
	case prod.Quantity <= 0:
		return nil, config.ErrorProductQuantity
	}

	productModel := product.Product{
		Id:       utility.RandomId(),
		Name:     prod.Name,
		Price:    prod.Price,
		Quantity: prod.Quantity,
	}

	err := ps.ProductModel.CreateProduct(ctx, productModel)
	if err != nil {
		return nil, err
	}

	// delete cache product
	_ = ps.Memory.Delete("products")

	return (*ProductResponse)(&productModel), nil
}

func (ps *ProductService) UpdateProduct(ctx context.Context, prod UpdateProduct) error {
	switch {
	case len(prod.Id) < 20:
		return errors.New("format id invalid")
	case len(prod.Name) < 2:
		return config.ErrorProductName
	case prod.Price == 0:
		return config.ErrorProductPrice
	case prod.Quantity <= 0:
		return config.ErrorProductQuantity
	}

	productServer := product.Product{
		Id:       prod.Id,
		Name:     prod.Name,
		Price:    prod.Price,
		Quantity: prod.Quantity,
	}

	err := ps.ProductModel.UpdateProduct(ctx, productServer)
	if err != nil {
		return err
	}

	// delete cache product
	_ = ps.Memory.Delete("products")

	return nil
}

func (ps *ProductService) DeleteProductById(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id not found")
	}

	err := ps.ProductModel.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}

	_ = ps.Memory.Delete("products")
	return nil
}

// Set cache in memory
func (ps *ProductService) SetCacheProducts(products []ProductResponse) error {
	if products == nil {
		return errors.New("products not found")
	}

	cacheProducts, err := ffjson.Marshal(products)
	if err != nil {
		return err
	}

	if err = ps.Memory.Set("products", cacheProducts); err != nil {
		return err
	}

	return nil
}
