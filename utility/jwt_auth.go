package utility

import (
	"errors"
	"time"

	"github.com/izzxx/Go-Restful-Api/config"
	"github.com/golang-jwt/jwt"
)

type JwtSchema struct {
	Name    string
	Email   string
	IsAdmin bool
	jwt.StandardClaims
}

func GenerateToken(admin bool, email, name string) (string, error) {
	jwtSchema := JwtSchema{
		Name:    name,
		Email:   email,
		IsAdmin: admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtSchema)

	token, err := jwtToken.SignedString([]byte(config.JwtSecretkey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(token string) (*JwtSchema, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtSchema{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecretkey), nil
	})
	if err != nil || !tokenClaims.Valid {
		return nil, err
	}

	schema, ok := tokenClaims.Claims.(*JwtSchema)
	if !ok {
		return nil, errors.New("failed to convert claims to jwtSchema")
	}

	return schema, nil
}
