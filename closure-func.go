package main

import (
	"fmt"
)

func main() {

	var minMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 5, 6, 7, 8, 12, 13, 99}
	var min, max = minMax(numbers)

	fmt.Printf("Data: %v \nMin: %v \nMax: %v \n", numbers, min, max)

}
