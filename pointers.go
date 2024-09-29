//Go dilinde pointer kullanarak bir integer dizisi için bellekte dinamik olarak yer ayırma

package main

import "fmt"

func main() {
	size := 5
	ptr := new(*int)
	arr := make([]int, size)
	*ptr = &arr[0]

	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}

	for i := 0; i < size; i++ {
		fmt.Println("arr[", i, "] =", arr[i])
	}

}