package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pizza:")
	input, _ := reader.ReadString('\n')

	fmt.Println(input)
	fmt.Printf("Type of rating is %T", input)
}
