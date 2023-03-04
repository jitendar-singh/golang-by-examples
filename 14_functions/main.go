package main

import "fmt"

func main() {
	greeter()
	fmt.Println(adder(10, 20))
	sum, msg := proadder(10, 20, 30, 40, 50)
	fmt.Println(msg, "-", sum)
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func proadder(values ...int) (int, string) {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum, "Total Sum:"
}
