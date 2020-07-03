package main

import (
	"fmt"
)

func main() {

	var buah = []string{"apel", "anggur", "melon", "semangka"}
	
	for i := 0; i < len(buah); i++{
		fmt.Printf("Index ke %d adalah buah %s \n", i, buah[i])
		
	}
}