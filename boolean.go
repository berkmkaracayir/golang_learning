package main

import (
"fmt"
"strconv"
)

func main() {
var numbers []int

    for {
    	var input string
    	fmt.Print("Enter a number or 'done' to finish: ")
    	fmt.Scanln(&input)

    	if input == "done" {
    		break
    	}

    	number, err := strconv.Atoi(input)
    	if err != nil {
    		fmt.Println("Invalid input")
    		continue
    	}

    	numbers = append(numbers, number)
    }

    sum := calculateSum(numbers)
    fmt.Println("The sum of the numbers is:", sum)

}

func calculateSum(numbers []int) int {
sum := 0

    for _, number := range numbers {
    	sum += number
    }

    return sum

}