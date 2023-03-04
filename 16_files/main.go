package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to go inside the file, and we will check is later"

	file, err := os.Create("./mygofile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./mygofile.txt")
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
