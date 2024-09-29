package main

import (
	"fmt"
)

func sayilariTopla(a, b int) int {
    return a + b
}

func main() {
	toplam := sayilariTopla(10, 20)

	fmt.Println("Toplam: ", toplam)
}