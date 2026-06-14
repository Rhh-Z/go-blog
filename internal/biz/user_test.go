package biz

import (
	"fmt"
	"testing"
)

func TestHashPassword(test *testing.T) {
	s := hashPassword("acc")
	fmt.Println(s)
}

func TestVerifyPassword(test *testing.T) {
	b1 := verifyPassword("$2a$10$J4nxvUuSAU0NrMTxjYAUxOmN/zRYErJc8IFn4SNFt1GesU8Ms9Jqe", "aaa")
	fmt.Println(b1)

	b2 := verifyPassword("$2a$10$J4nxvUuSAU0NrMTxjYAUxOmN/zRYErJc8IFn4SNFt1GesU8Ms9Jqe", "acc")
	fmt.Println(b2)
}
