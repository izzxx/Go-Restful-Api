package config

import "errors"

var (
	ErrorEmail           = errors.New("an error occured while entering the email")
	ErrorPassword        = errors.New("password letters cannot be less then 8")
	ErrorProductName     = errors.New("product name letters cannot be less than 2")
	ErrorProductPrice    = errors.New("price cannot be less than 0")
	ErrorProductQuantity = errors.New("quantity cannot be zero or less than 0")
)
