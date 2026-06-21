package auth

import (
	"fmt"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	token := GenerateToken("secret", "12312")
	fmt.Println(token)
}
