package utility

import (
	"log"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken(false, "example@email.mail", "example name")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(token)
}

func TestValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoiZXhhbXBsZSBuYW1lIiwiRW1haWwiOiJleGFtcGxlQGVtYWlsLm1haWwiLCJleHAiOjE2NTM5MjYzMjZ9.zDYbxKnLn5cCHbsaYJUKmubwSU509L33iiuvmWKoBIw"
	schema, err := ValidateToken(token)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("name: %s\nemail: %s\n", schema.Name, schema.Email)
}
