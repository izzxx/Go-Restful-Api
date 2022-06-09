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

	_, err = tx.Exec(ctx, "INSERT INTO products(id, name, price, quantity) VALUES($1, $2, $3, $4)", prod.Id, prod.Name, prod.Price, prod.Quantity)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
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

	var product Product

	err = tx.QueryRow(ctx, "SELECT id, name, price, quantity, created_at FROM products WHERE id = $1", id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Quantity,
		&product.Created_At,
	)

	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
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

	rows, err := tx.Query(ctx, "SELECT id, name, price, quantity, created_at FROM products ORDER BY name ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var product Product

		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity, &product.Created_At)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	err = tx.Commit(ctx)
	if err != nil {
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

	ct, err := tx.Exec(ctx, "UPDATE products SET name = $1, price = $2, quantity = $3 WHERE id = $4", &prod.Name, &prod.Price, &prod.Quantity, &prod.Id)
	if err != nil || ct.RowsAffected() == 0 {
		return errors.New("failed to update product")
	}

	err = tx.Commit(ctx)
	if err != nil {
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

	ct, err := tx.Exec(ctx, "DELETE FROM products WHERE id = $1", id)
	if err != nil || ct.RowsAffected() == 0 {
		return errors.New("product not found")
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
