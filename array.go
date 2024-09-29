package main

import (
	"fmt"
)

func main() {
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50

	for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
}