package main

import "fmt"

func main() {
    var rows, columns, i, j int

    var mat1 [10][10]int
    var mat2 [10][10]int
    var multiplicationmat [10][10]int

    fmt.Print("Satır ve Sütun Sayılarını Girin = ")
    fmt.Scan(&rows, &columns)

    fmt.Print("Birinci Matrisin Elemanlarını Girin = ")
    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            fmt.Scan(&mat1[i][j])
        }
    }

    fmt.Print("İkinci Matrisin Elemanlarını Girin = ")
    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            fmt.Scan(&mat2[i][j])
        }
    }

    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            multiplicationmat[i][j] = 0
            for k := 0; k < columns; k++ {
                multiplicationmat[i][j] += mat1[i][k] * mat2[k][j]
            }
        }
    }
    fmt.Println("-- Matris Çarpımının Sonucu --")
    for i = 0; i < rows; i++ {
        for j = 0; j < columns; j++ {
            fmt.Print(multiplicationmat[i][j], "\t")
        }
        fmt.Println()
    }
}