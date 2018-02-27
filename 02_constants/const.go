package main

import "fmt"

const p string = "death & taxes"

var p = "new value"

func main() {
	const q = 42

	fmt.Println("p -", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchanging value
