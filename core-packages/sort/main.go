package main

import (
	"fmt"
	"sort"
)

func main() {
	vars_ints := []int{5, 2, 0, 3, 4, 9, 6}
	vars_strings := []string{"a", "b", "c", "z", "x"}
	
	sort.Ints(vars_ints)
	sort.Strings(vars_strings)

	fmt.Println(vars_ints)
	fmt.Println(vars_strings)
}
