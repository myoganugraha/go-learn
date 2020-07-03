package main

import (
	"fmt"
)

func main() {

	showMessage()
	fmt.Println(showMessageWithStringReturn())
	fmt.Println("this is value from showMessageWithIntReturn()", showMessageWithIntReturn())
	fmt.Println("this is value from showMessageWithIntReturnAndParam() with x = 10, and y = 2,", showMessageWithIntReturnAndParam(10,2))
}

func showMessage() {
	fmt.Println("this is a message from showMessage() function")
}
func showMessageWithStringReturn() string {
	return "this is a message from showMessageWithStringReturn() function"
}

func showMessageWithIntReturn() int {
	return 10
}

func showMessageWithIntReturnAndParam(x int, y int) int {
	return  x + y
}