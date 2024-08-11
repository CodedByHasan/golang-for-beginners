package main

import "fmt"

func modify(str *string) {
	// Dereference and then change value
	*str = "world"
}

func modifySlice(slice []int) {
	slice[0] = 100
}

func modifyMap(m map[string]int) {
	m["K"] = 75
}

func main() {
	strA := "hello"
	fmt.Printf("BEFORE modify(): %s\n", strA)

	modify(&strA) // pass by reference
	fmt.Printf("AFTER modify(): %s\n", strA)

	slice := []int{10, 20, 30}
	fmt.Println("BEFORE modifySlice():", slice)

	// slices are passed by reference by default
	modifySlice(slice)
	fmt.Println("AFTER modifySlice():", slice)

	asciiCodes := make(map[string]int)
	asciiCodes["A"] = 65
	asciiCodes["F"] = 70
	fmt.Println("BEFORE modifyMap():", asciiCodes)

	// maps are passed by reference by default
	modifyMap(asciiCodes)
	fmt.Println("AFTER modifyMap():", asciiCodes)

}
