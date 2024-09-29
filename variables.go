package main
import "fmt"

func main() {

    // `var` 1 veya 1 den fazla değişkeni tanımlar
    var a = "başlangıç değeri "
    fmt.Println(a)

    // 1 seferde 1 den fazla değişken tanımlayabilirsiniz.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go başlatılan değişkenlerin tipini algılar
    var d = true
    fmt.Println(d)

    // Değer atanmadan tanımlanan değişkenler 0 değerlidir
    var e int
    fmt.Println(e)

    // The `:=` syntax bir değişkene "initial" değeri vermek için kullanılır
    var f string = "elma"
    fmt.Println(f)

	var g int = 10
	fmt.Println(g)

}
