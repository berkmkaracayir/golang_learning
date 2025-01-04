/*Bu kod, bir sayıların listesinin ortalamasını hesaplar. 
Ancak, listenin tamamını calculateAverage fonksiyonuna geçirirken kopyalar. 
Liste büyükse bu işlem verimsiz olabilir. 
Listeyi kopyalamamak için, liste için bir işaretçi kullanabilirsiniz:*/

package main

import "fmt"

func main() {
	numbers := []float64{1, 2, 3, 4, 5}
	average := calculateAverage(&numbers)
	fmt.Printf("The average of the numbers %v is %v.\n", numbers, average)
}

func calculateAverage(numbers *[]float64) float64 {
	var sum float64
	for _, number := range *numbers {
		sum += number
	}
	return sum / float64(len(*numbers))
}