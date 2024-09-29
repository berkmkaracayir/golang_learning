// Kanallar, gorutinler arasında iletişim kurmak için kullanılır.
// Kanallar, paylaşılan kaynaklara erişimi senkronize etmek veya programın farklı bölümleri arasında veri aktarmak için kullanılabilir.
package main
import "fmt"
func sum(values []int, result chan int) {
    sum := 0
    for _, v := range values {
        sum += v
    }
    result <- sum // toplamı channel a gönder
// Bir değeri kanaldan göndermek için <- operatörü kullanılabilir
// x := <-ch // kanaldan bir değer al ve x değişkenine kaydet

}

func main() {
    values := []int{1, 2, 3, 4, 5}

    result := make(chan int) //ch := make(chan int)  bir tamsayı kanalı oluştur
    go sum(values[:len(values)/2], result) // sum ı gorutinde çalıştır
    go sum(values[len(values)/2:], result) // sum ı gorutinde çalıştır

    x, y := <-result, <-result // kanaldan toplamları al
    fmt.Println(x + y) 
}