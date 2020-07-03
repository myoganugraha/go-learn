package main

import (
	"fmt"
)

func main() {
	nilai := 21

	if nilai % 2 == 0 {
		fmt.Printf("Nilai %d merupakan bilangan genap \n", nilai)
	} else {
		fmt.Printf("Nilai %d merupakan bilangan ganjil \n", nilai)
	}
}