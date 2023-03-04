package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("Index is %v & value is %v\n", index, day)
	}

	rougevalue := 1
	for rougevalue < 10 {
		fmt.Println("value is ", rougevalue)
		rougevalue++
		if rougevalue == 5 {
			break //continue
		}
	}
}
