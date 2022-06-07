package user

import (
	"context"
	"errors"
	"net/mail"

	"github.com/google/uuid"
	"github.com/izzxx/Go-Restful-Api/config"
	"github.com/izzxx/Go-Restful-Api/model/user"
	"github.com/izzxx/Go-Restful-Api/utility"
)

type UserService struct {
	UserRepository user.UserRepository
}

func (us *UserService) Register(ctx context.Context, guestRegister UserRegister) (*UserResponse, error) {
	// Email validation
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

	err = guest.HashPassword()
	if err != nil {
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
	// Email validation
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

	err = user.ComparePassword(guestLogin.Password)
	if err != nil {
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
	// Email validation
	_, err := mail.ParseAddress(userUpdate.Email)

	switch {
	case err != nil:
		return config.ErrorEmail
	case len(userUpdate.PastPassword) < 8:
		return config.ErrorPassword
	case len(userUpdate.NewPassword) < 8:
		return config.ErrorPassword
	}

	user, err := us.UserRepository.GetUserByEmail(ctx, userUpdate.Email)
	if err != nil {
		return err
	}

	// Compare user password
	err = user.ComparePassword(userUpdate.PastPassword)
	if err != nil {
		return err
	}

	// Reset old password with new password
	user.Password = userUpdate.NewPassword

	// Hash new password
	err = user.HashPassword()
	if err != nil {
		return err
	}

	// Change old password to new password
	err = us.UserRepository.UpdatePasswordUser(ctx, userUpdate.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}
