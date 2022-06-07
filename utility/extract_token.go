package utility

import (
	"errors"
	"strings"
)

func ExtractToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("token not found")
	}

	match := strings.Contains(token, "Bearer ")
	if !match {
		return "", errors.New("wrong format token")
	}

	tokenSlice := strings.Split(token, " ")
	if len(tokenSlice) != 2 {
		return "", errors.New("wrong format token")
	}

	return tokenSlice[1], nil
}
