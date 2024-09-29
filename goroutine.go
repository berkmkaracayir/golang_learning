//çalışma zamanı tarafından yönetilen hafif iş parçacıkları

package main
import "fmt"

func sayHello() {
    fmt.Println("Hello, world!")
}

func main() {
    go sayHello() // start a new goroutine
    fmt.Println("Main function")
}