package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname, lname string) string {

	return fmt.Sprint(fname, lname)
}

// test on new branch from office
// add a comment on Sunday 25 number 2
