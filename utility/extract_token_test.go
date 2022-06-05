package utility

import (
	"log"
	"testing"
)

func TestExtractTokenValid(t *testing.T) {
	tokenValid := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiZXhhbXBsZSBuYW1lIiwiRW1haWwiOiJleGFtcGxlQGVtYWlsLm1haWwiLCJleHAiOjE2NTM5MjYzMjZ9.zDYbxKnLn5cCHbsaYJUKmubwSU509L33iiuvmWKoBIw"

	extractToken, err := ExtractToken(tokenValid)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(extractToken)
}

func TestExtractTokenInvalid(t *testing.T) {
	tokenInvalid := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiZXhhbXBsZSBuYW1lIiwiRW1haWwiOiJleGFtcGxlQGVtYWlsLm1haWwiLCJleHAiOjE2NTM5MjYzMjZ9.zDYbxKnLn5cCHbsaYJUKmubwSU509L33iiuvmWKoBIw"

	extractToken, err := ExtractToken(tokenInvalid)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(extractToken)
}
