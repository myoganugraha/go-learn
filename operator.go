package main

import (
	"fmt"
)

func main() {
	firstValue := 10
	secondValue := 23

	fmt.Printf("Penjumlahan %d \n", firstValue + secondValue)
	fmt.Printf("Pengurangan %d \n", firstValue - secondValue)
	fmt.Printf("Perkalian %d \n", firstValue * secondValue)
	fmt.Printf("Pembagian %.f \n", float32(secondValue / firstValue))
	fmt.Printf("Hasil bagi %.f \n", float64(secondValue % firstValue))
}