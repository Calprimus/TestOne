package main

import "fmt"

func main() {
	var smallNum int
	var largeNum int

	fmt.Print("Please enter a large number: ")
	fmt.Scan(&smallNum)
	fmt.Print("Please enter a small number: ")
	fmt.Scan(&largeNum)
	var remNum = largeNum % smallNum
	fmt.Println("And the result is", remNum)
}
