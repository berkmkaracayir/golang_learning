//go mod init main.go
//go mod tidy
//go test

package main

import (
	"fmt"
)

func main() {
	result := sum(2, 3)
	fmt.Println("Result:", result)
}

func sum(a, b int) int {
	return a + b
}