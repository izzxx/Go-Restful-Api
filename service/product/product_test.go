package product

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/izzxx/Go-Restful-Api/model/product"
	"github.com/jackc/pgx/v4/pgxpool"
)

var db *pgxpool.Pool
var memory *bigcache.BigCache
var ctx context.Context
var service *ProductService

func TestMain(m *testing.M) {
	cntx, cf := context.WithDeadline(context.Background(), time.Now().Add(3*time.Minute))
	defer cf()

	ctx = cntx

	coon, err := pgxpool.ParseConfig(os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}

	db, err = pgxpool.ConnectConfig(ctx, coon)
	if err != nil {
		log.Fatal(err)
	}

	memory, err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Minute))
	if err != nil {
		log.Fatal(err)
	}

	service = &ProductService{
		ProductModel: product.ProductModel{Db: db},
		Memory:       memory,
	}

	os.Exit(m.Run())
}

func TestGetAllProductFromDb(t *testing.T) {
	productsResponse, err := service.GetAllProductsFromDb(ctx)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(productsResponse)
}

func TestGetProductByIdFromCache(t *testing.T) {
	// After fetching all products, cache will be set
	_, err := service.GetAllProductsFromDb(ctx)
	if err != nil {
		t.Error(err)
	}

	product, err := service.GetProductIdFromCache("FsA6QNjP5M7GHqETPz4b")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(product)
}

func TestStoreNewProduct(t *testing.T) {
	newProduct := StoreProduct{
		Name:     "Iphone 12 Pro Max",
		Price:    13.900099,
		Quantity: 100,
	}

	product, err := service.StoreProduct(ctx, newProduct)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(product.Id)
	fmt.Println(product)
}

func TestUpdateProduct(t *testing.T) {
	product := UpdateProduct{
		Id:       "FsA6QNjP5M7GHqETPz4b",
		Name:     "Iphone 12 Pro Max",
		Price:    13.900099,
		Quantity: 149,
	}

	err := service.UpdateProduct(ctx, product)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("success")
}

func TestDeleteProductById(t *testing.T) {
	productId := "FsA6QNjP5M7GHqETPz4b"

	err := service.DeleteProductById(ctx, productId)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("success")
}
