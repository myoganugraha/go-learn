package main

import (
	"fmt"
)

func main() {

	var buah = [4]string{"apel", "anggur", "melon", "semangka"}

	fmt.Println("Jumlah buah", len(buah))
	fmt.Println("Array pertama", buah[0])
	fmt.Println("isi array", buah)

	buah[1] = "anjay"
	fmt.Println("isi buah setelah array index kedua diubah", buah)
}