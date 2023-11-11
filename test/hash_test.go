package test

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	bytes, _ := bcrypt.GenerateFromPassword([]byte("12341234"), 14)
	fmt.Println(string(bytes))

}
