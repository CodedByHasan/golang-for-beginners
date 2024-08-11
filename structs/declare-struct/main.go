package main

import "fmt"

type Student struct {
	name   string
	id     int
	marks  []int
	grades map[string]int
}

func main() {
	var harryPotter Student
	fmt.Printf("%+v\n", harryPotter)
}
