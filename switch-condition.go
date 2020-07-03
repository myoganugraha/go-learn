package main

import (
	"fmt"
)

func main() {
	nilai := 2

	switch {
	case nilai ==10:
		fmt.Printf("Perfecto dengan nilai %d \n", nilai)
	case nilai > 7:
		fmt.Printf("Lulus aja dengan nilai %d \n", nilai)
	default:
		fmt.Printf("Gak lulus, nilai %d \n", nilai )
	}
}