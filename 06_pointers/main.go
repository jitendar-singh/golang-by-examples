package main

import "fmt"

func main() {
	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of actual pointer:", ptr)
	fmt.Println("value of actual pointer:", *ptr)

	*ptr += 2
	fmt.Println("New Value at pointer:", *ptr)
	fmt.Println("New Value at number:", myNumber)

	myNumber += 2
	fmt.Println("New Value at pointer:", *ptr)
	fmt.Println("New Value at number:", myNumber)
}
