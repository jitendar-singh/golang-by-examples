package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Java Script"
	languages["py"] = "Python"
	languages["rb"] = "Ruby"

	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("The .%v extension stands for %v\n ", key, value)
	}
}
