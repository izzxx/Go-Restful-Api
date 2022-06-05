package user

import (
	"context"
	"errors"
	"net/mail"

	"github.com/izzxx/Go-Restful-Api/config"
	"github.com/izzxx/Go-Restful-Api/model/user"
	"github.com/izzxx/Go-Restful-Api/utility"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository user.UserRepository
}

func (us *UserService) Register(ctx context.Context, guestRegister UserRegister) (*UserResponse, error) {
	_, err := mail.ParseAddress(guestRegister.Email)

	switch {
	case err != nil:
		return nil, config.ErrorEmail
	case len(guestRegister.Username) < 2:
		return nil, errors.New("username letters cannot be less than 2")
	case len(guestRegister.Password) < 8:
		return nil, config.ErrorPassword
	}

	guest := user.User{
		Id:       uuid.NewString(),
		Name:     guestRegister.Username,
		Email:    guestRegister.Email,
		Password: guestRegister.Password,
		IsAdmin:  guestRegister.IsAdmin,
	}

	if err = guest.HashPassword(); err != nil {
		return nil, err
	}

	id, err := us.UserRepository.CreateUser(ctx, guest)
	if err != nil {
		return nil, err
	}

	// generate jwt token after success register
	token, err := utility.GenerateToken(guest.IsAdmin, guest.Email, guestRegister.Username)
	if err != nil {
		return nil, err
	}

	responseApi := UserResponse{
		Id:    id,
		Email: guest.Email,
		Token: token,
	}

	return &responseApi, nil
}

func (us *UserService) Login(ctx context.Context, guestLogin UserLogin) (*UserResponse, error) {
	_, err := mail.ParseAddress(guestLogin.Email)

	switch {
	case err != nil:
		return nil, config.ErrorEmail
	case len(guestLogin.Password) < 8:
		return nil, config.ErrorPassword
	}

	user, err := us.UserRepository.GetUserByEmail(ctx, guestLogin.Email)
	if err != nil {
		return nil, err
	}

	if err = user.ComparePassword(guestLogin.Password); err != nil {
		return nil, err
	}

	// generate jwt token after success login
	token, err := utility.GenerateToken(user.IsAdmin, user.Email, user.Name)
	if err != nil {
		return nil, err
	}

	responseApi := UserResponse{
		Id:    user.Id,
		Email: user.Email,
		Token: token,
	}

	return &responseApi, nil
}

func (us *UserService) UpdatePassword(ctx context.Context, userUpdate UserUpdatePassword) error {
	_, err := mail.ParseAddress(userUpdate.Email)

	switch {
	case err != nil:
		return config.ErrorEmail
	case len(userUpdate.PastPassword) < 8:
		return config.ErrorPassword
	case len(userUpdate.NewPassword) < 8:
		return config.ErrorPassword
	default:
	}

	var user = user.User{
		Password: userUpdate.PastPassword,
	}

	if err = user.HashPassword(); err != nil {
		return err
	}

	if err = us.UserRepository.UpdatePasswordUser(ctx, userUpdate.Email, user.Password, userUpdate.NewPassword); err != nil {
		return err
	}

	return nil
}
