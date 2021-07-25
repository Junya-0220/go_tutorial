package main

import (
	"fmt"
)

func main() {
	f := 1.11
	i := int(f)
	fmt.Printf("%T %v\n", i, i)

	s := []int{1, 2, 5, 6, 2, 3, 1}
	fmt.Println(s[:])

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"messi": 30,
	}
	fmt.Printf("%T, %v", m, m)
}
