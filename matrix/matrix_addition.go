package main

import "fmt"

func main() {
    var i, j, rows, columns int

    var firstMat [10][10]int
    var secondMat [10][10]int
	var sumMat [10][10]int

    fmt.Print("Matrisin satır ve sütun sayısını girin = ")
    fmt.Scan(&rows, &columns)

    fmt.Println("Birinci Matris Elemanlarını Girin = ")
    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            fmt.Scan(&firstMat[i][j])
        }
    }
    fmt.Println("İkinci Matris Elemanlarını Girin = ")
    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            fmt.Scan(&secondMat[i][j])
        }
    }
  	for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            sumMat[i][j]=firstMat[i][j]+secondMat[i][j]
        }
    }
	for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
        	    fmt.Print(sumMat[i][j], "  ")
        }
        fmt.Println()        
	}    	
}