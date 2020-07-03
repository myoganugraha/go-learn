package main

import (
	"fmt"
)

func main() {

	var foodPrice = map[string]int{"fried_chicken": 18000, "fried_rice": 13000}
	fmt.Println("fried chicken price: Rp.", foodPrice["fried_chicken"])
}