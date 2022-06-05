package product

import (
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/izzxx/Go-Restful-Api/model/product"
)

type ProductResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Price      float64   `json:"price"`
	Quantity   uint16    `json:"quantity"`
	Created_At time.Time `json:"created_at"`
}

type StoreProduct struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint16  `json:"quantity"`
}

type UpdateProduct struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity uint16  `json:"quantity"`
}

type ProductService struct {
	ProductModel product.ProductModel
	Memory       *bigcache.BigCache
}
