package main

import "fmt"

type Student struct {
	name string
	id   int
}

func main() {
	harryPotter := new(Student)
	fmt.Printf("%+v\n", harryPotter)

	ronWeasly := Student{
		name: "Ron",
		id:   1,
	}

	fmt.Printf("%+v\n", ronWeasly)

	// omitting field names (not recommended)
	tomRiddle := Student{"Tom", 2}
	fmt.Printf("%+v\n", tomRiddle)

}
