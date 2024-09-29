//Go, fonksiyonlara ... sözdizimi kullanarak değişken sayıda argüman geçme olanağı tanır.

package main

import (
	"fmt"
)

func ortalamaHesapla(sayilar ...float64) float64 {
    toplam := 0.0
    for _, sayi := range sayilar {
        toplam += sayi
    }
    return toplam / float64(len(sayilar))
}

func main() {
	ortalama := ortalamaHesapla(10.5, 7.2, 4.8, 9.1, 6.6)

	fmt.Println("Ortalama: ", ortalama)
}