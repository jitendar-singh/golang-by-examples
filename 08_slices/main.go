package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{}
	fmt.Printf("Type of fruitList: %T\n", fruitList)

	var fruits = []string{"Apple", "Oranges", "Peach"}
	fmt.Println("Fruits slice:", fruits)
	fruits = append(fruits, "Water melon", "Papaya", "Guava")
	fmt.Println("Fruits post appending:", fruits)

	fruits = append(fruits[1:])
	fmt.Println("Appending fruits[1:] to fruits:", fruits)

	fruits = append(fruits[1:3])
	fmt.Println("Appending fruits[1:3] to fruits:", fruits)

	fruits = append(fruits[:4])
	fmt.Println("Appending fruits[:4] to fruits:", fruits)

	highScore := make([]int, 4)
	highScore = append(highScore, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(highScore)
	highScore[0] = 11
	highScore[1] = 22
	highScore[2] = 33
	highScore[3] = 44
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	var courses = []string{"javascript", "go", "python", "react", "ruby"}
	fmt.Println(courses)
	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
