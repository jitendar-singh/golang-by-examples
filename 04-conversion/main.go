package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter your rating")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating us:", input)

	fmt.Println("Adding 1 to your rating")
	newInput, error := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(newInput + 1)
	}
}
