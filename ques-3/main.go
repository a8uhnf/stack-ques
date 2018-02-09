package main

import (
	"fmt"
	"context"
)

type User struct {
	UserID              int
	Email               string
	FirstName           string
	LastName            string
	PasswordHash        string
}

func main() {
	fmt.Println("------------------")
	ctx := context.Background()

	user := &User{"Hello", "abc@gmail.com", "abc", "def", "12oajoajbfabf82y914919417"}
}