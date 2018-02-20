package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 25; i++ {
		fmt.Printf("%b - %d - %x \n", i, i, i)
	}
}

func mainNew() {
	x := 17
	fmt.Println(x)

}
