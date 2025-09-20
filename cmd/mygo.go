package main

import (
	"fmt"

	uuid "github.com/google/uuid"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("Generated UUID:", uuid.New().String())
}
