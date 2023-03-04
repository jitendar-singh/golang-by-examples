package main

import "fmt"

func main() {
	var veggyList [4]string
	veggyList[0] = "Potato"
	veggyList[1] = "Beans"
	veggyList[2] = "Carrot"

	fmt.Println("Vegetable List is: ", veggyList)
	fmt.Println("Vegetable list length is:", len(veggyList))

	newVeggyList := [4]string{
		"Radish",
		"Cucumber",
		"Ladyfinger",
	}
	fmt.Println(newVeggyList)
}
