package main

import (
	"fmt"
)

func bol(a, b int) (int, int) {
    bolum := a / b
    kalan := a % b
    return bolum, kalan
}

func main() {
	b, k := bol(10, 3)

	fmt.Println("Bölüm: ", b)
	fmt.Println("Kalan: ", k)
}