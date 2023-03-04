package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	jitsingh := User{
		"Jitendar Singh",
		"jitsingh@redhat.com",
		33,
		true,
	}

	fmt.Println(jitsingh)
	fmt.Printf("%+v\n", jitsingh)
	fmt.Println(jitsingh.Email)
}
