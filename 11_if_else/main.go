package main

import "fmt"

func main() {
	count := make([]int, 3)
	count[0] = 23
	count[1] = 2
	count[2] = 10

	for value := range count {
		if count[value] < 10 {
			fmt.Println("count is less than 10")
		} else if count[value] > 10 {
			fmt.Println("count is greater than 10")
		} else {
			fmt.Println("count is exactly 10")
		}
	}

	if num := 3; num > 10 {
		fmt.Println("Number is greater than 10")
	} else {
		fmt.Println("Number is less than or equal to 10")
	}

}
