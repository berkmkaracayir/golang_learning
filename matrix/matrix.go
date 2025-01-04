package main

import "fmt"

func main() {
    var i, j, rows, columns int

    var matrix [10][10]int
    var transposeMatrix [10][10]int

    fmt.Print("Matrisin satır ve sütun sayısını girin = ")
    fmt.Scan(&rows, &columns)

    fmt.Println("Transpoze etmek için Matris Elemanlarını Girin = ")
    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            fmt.Scan(&matrix[i][j])
        }
    }
    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            transposeMatrix[j][i] = matrix[i][j]
        }
    }
    fmt.Println("--- Transpoze Matris Elemanları ---")
    for i = 0; i < columns; i++ {
        for j = 0; j < rows; j++ {
            fmt.Print(transposeMatrix[i][j], "  ")
        }
        fmt.Println()
    }
}