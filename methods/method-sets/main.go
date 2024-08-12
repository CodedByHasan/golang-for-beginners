package main

import "fmt"

type Student struct {
	name   string
	grades []int
}

func (s *Student) displayName() {
	fmt.Println(s.name)
}

func (s *Student) calculatePercentage() float64 {
	sum := 0
	for _, value := range s.grades {
		sum += value
	}
	return float64(sum*100) / float64(len(s.grades)*100)
}

func main() {
	student := Student{name: "Peter Parker", grades: []int{70, 80, 90}}

	// displayName and calculatePercentage are method set of Student
	student.displayName()
	fmt.Printf("%.2f%%\n", student.calculatePercentage())
}
