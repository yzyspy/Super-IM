package jwt

import (
	"fmt"
	"testing"
)

func TestJwtGenerateToken(t *testing.T) {
	token, err := GenerateJWT("123", "testuser")
	if err != nil {
		t.Errorf("GenerateJWT() error = %v", err)
		return
	}
	fmt.Println(token)

	jwt, err := ParseJWT(token)
	fmt.Println(jwt)
}
