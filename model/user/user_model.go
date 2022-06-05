package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) HashPassword() error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)

	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

type UserRepository struct {
	Db *pgxpool.Pool
}

func (ur *UserRepository) CreateUser(ctx context.Context, user User) (string, error) {
	coon, err := ur.Db.Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	query := `INSERT INTO users(id, name, email, password, isadmin) VALUES($1, $2, $3, $4, $5) RETURNING id`
	if err = tx.QueryRow(ctx, query, user.Id, user.Name, user.Email, user.Password, user.IsAdmin).Scan(&user.Id); err != nil {
		return "", err
	}

	if err = tx.Commit(ctx); err != nil {
		return "", err
	}

	return user.Id, nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	coon, err := ur.Db.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	var usr User

	query := `SELECT id, name, email, password, isadmin FROM users WHERE email = $1`
	if err = tx.QueryRow(ctx, query, email).Scan(&usr.Id, &usr.Name, &usr.Email, &usr.Password, &usr.IsAdmin); err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &usr, nil
}

func (ur *UserRepository) UpdatePasswordUser(ctx context.Context, email, pastPassword, newPassword string) error {
	coon, err := ur.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `UPDATE users SET password = $1 WHERE email = $2 AND password = $3`
	if ct, err := tx.Exec(ctx, query, newPassword, email, pastPassword); err != nil || ct.RowsAffected() == 0 {
		return errors.New("failed to update user")
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(ctx context.Context, email string) error {
	coon, err := ur.Db.Acquire(ctx)
	if err != nil {
		return err
	}
	defer coon.Release()

	tx, err := coon.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	query := `DELETE FROM users WHERE email = $1`
	if ct, err := tx.Exec(ctx, query, email); err != nil || ct.RowsAffected() == 0 {
		return errors.New("failed to delete user")
	}

	if err = tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}
