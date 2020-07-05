package main

import (
	"fmt"
)

func main() {

	var median = medianFunc(2, 4, 5, 6, 7, 8, 3, 4)
	fmt.Printf("Rata - rata %.2f \n", median)

	// fmt.Println(msg)

}

func medianFunc(numbers ...int) float64 {
	var total int

	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
