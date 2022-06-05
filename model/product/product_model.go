package product

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
)

type ProductModel struct {
	Db *pgxpool.Pool
}

func (pm *ProductModel) CreateProduct(ctx context.Context, prod Product) error {
	coon, err := pm.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := "INSERT INTO products(id, name, price, quantity) VALUES($1, $2, $3, $4)"
	if _, err = tx.Exec(ctx, query, prod.Id, prod.Name, prod.Price, prod.Quantity); err != nil {
		return err
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (pm *ProductModel) FindById(ctx context.Context, id string) (*Product, error) {
	coon, err := pm.Db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var product = Product{}

	query := "SELECT id, name, price, quantity, created_at FROM products WHERE id = $1"
	if err = tx.QueryRow(ctx, query, id).Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Created_At); err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &product, nil
}

func (pm *ProductModel) FindAllProduct(ctx context.Context) ([]Product, error) {
	coon, err := pm.Db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	// nil slice
	var products []Product

	query := "SELECT id, name, price, quantity, created_at FROM products ORDER BY name ASC"
	rows, err := tx.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product = Product{}

		if err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Created_At); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return products, nil
}

func (pm *ProductModel) UpdateProduct(ctx context.Context, prod Product) error {
	coon, err := pm.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := "UPDATE products SET name = $1, price = $2, quantity = $3 WHERE id = $4"
	ct, err := tx.Exec(ctx, query, &prod.Name, &prod.Price, &prod.Quantity, &prod.Id)
	if err != nil || ct.RowsAffected() == 0 {
		return errors.New("failed to update product")
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (pm *ProductModel) DeleteProduct(ctx context.Context, id string) error {
	coon, err := pm.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := "DELETE FROM products WHERE id = $1"
	if ct, err := tx.Exec(ctx, query, id); err != nil || ct.RowsAffected() == 0 {
		return errors.New("product not found")
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
