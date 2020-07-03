package main

import (
	"fmt"
)

func main() {
	nilai := 8

	if nilai == 10 {
		fmt.Printf("Perfecto dengan nilai %d \n", nilai)
	} else if nilai > 7 {
		fmt.Printf("Lulus aja dengan nilai %d \n", nilai)
	} else {
		fmt.Printf("Gak lulus, nilai %d \n", nilai )
	}
}