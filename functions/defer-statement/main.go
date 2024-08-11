package main

import "fmt"

func printName(name string) {
	fmt.Println(name)
}

func printEmployeeNo(employeeNo int) {
	fmt.Println(employeeNo)
}

func printAddress(address string) {
	fmt.Println(address)
}

// A defer statement delays the execution of a function until
// surrounding function returns.
func main() {
	printName("Wu-Tang Clan")

	// The address will be displayed before employee no. due to defer
	defer printEmployeeNo(23)

	printAddress("Staten Island")
}
