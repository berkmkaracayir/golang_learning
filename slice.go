//Slice yapısı, sabit bir boyutu olmayan, dinamik bir dizi yapısını ifade eder.

package main

import (
	"fmt"
)

func main() {
	s := make([]int, 3) // 3 integer değer içeren bir slice oluşturur make fonksiyonu kullanılarak tanımlanırlar
	s[0] = 1            // ilk sayı 1
	s[1] = 2            // 2
	s[2] = 3            // 3 e eşitle
	s = append(s, 4)    // slice a yeni eleman ekler append fonksiyonu kullanılarak boyutları değiştirilebilir
	for i := 0; i < len(s); i++{
	fmt.Println(s[i])
	}
}