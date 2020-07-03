package main

import (
	"fmt"
)

func main() {

	var buah = []string{"apel", "anggur", "melon", "semangka"}

	fmt.Println("Jumlah buah", len(buah))
	fmt.Println("Array index pertama", buah[0])
	fmt.Println("isi array", buah)
	fmt.Println()

	buah[1] = "anjay"
	fmt.Println("isi buah setelah array index kedua diubah", buah)
	fmt.Println()

	buah = append(buah, "duren") 
	fmt.Println("Jumlah buah setelah ditambah", len(buah))
	fmt.Println("buah yang baru ditambahkan adalah ", buah[len(buah) - 1])
}