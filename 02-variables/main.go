package main

import "fmt"

const LoginToken string = "adwjfvjwsbdjkbqef"

func main() {
	var username string = "sunnyconcise"
	fmt.Println(username)
	fmt.Printf("The variable is of type: %T \n", username)

	var age uint8 = 255
	fmt.Println(age)
	fmt.Printf("The variable is of type: %T \n", age)

	var salary int = 2553577822222
	fmt.Println(salary)
	fmt.Printf("The variable is of type: %T \n", salary)

	var ctc float32 = 1350000.75268222
	fmt.Println(ctc)
	fmt.Printf("The variable is of type: %T \n", ctc)

	var isPresent bool
	fmt.Println(isPresent)
	fmt.Printf("The variable is of type: %T \n", isPresent)

	fmt.Println(LoginToken)
	fmt.Printf("The variable is of type: %T \n", LoginToken)

}
