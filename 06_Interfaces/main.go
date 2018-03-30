package main

import (
	"fmt"
	"sort"
)

// (1)
// type people []string
// studyGroup := people{"Zeno", "John", "Al", "Jenny"}

// (2)
// s := []string{"Zeno", "John", "Al", "Jenny"}

// (3)
// n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

// Also sort the above in reverse order
type people []string

// Ascending
// func main() {
// 	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
// 	sort.Sort(sort.StringSlice(studyGroup))
// 	fmt.Println(studyGroup)

// 	s := []string{"Zeno", "John", "Al", "Jenny"}
// 	sort.Sort(sort.StringSlice(s))
// 	fmt.Println(s)

// 	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
// 	sort.Sort(sort.IntSlice(n))
// 	fmt.Println(n)
// }

// Reverse
// func main() {
// 	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
// 	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
// 	fmt.Println(studyGroup)

// 	s := []string{"Zeno", "John", "Al", "Jenny"}
// 	sort.Sort(sort.Reverse(sort.StringSlice(s)))
// 	fmt.Println(s)

// 	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
// 	sort.Sort(sort.Reverse(sort.IntSlice(n)))
// 	fmt.Println(n)
// }

// func (p people) Len() int           { return len(p) }
// func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
// func (p people) Less(i, j int) bool { return p[i] < p[j] }

// func main() {
// 	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
// 	fmt.Println(studyGroup)
// 	sort.Sort(studyGroup)
// 	fmt.Println(studyGroup)
//}

func main() {
	// studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	// sort.Sort(sort.StringSlice(studyGroup))
	// fmt.Println(studyGroup)

	// s := []string{"Zeno", "John", "Al", "Jenny"}
	// sort.Sort(sort.StringSlice(s))
	// fmt.Println(s)

	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println(n)
}
