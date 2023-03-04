package main

import "fmt"

func main() {
	jitsingh := User{
		"Jitendar Singh",
		"jitsingh@redhat.com",
		true,
		33,
	}
	fmt.Println(jitsingh.Name)
	fmt.Println(jitsingh.Email)
	jitsingh.GetStatus()
	jitsingh.NewEmail()
	fmt.Println(jitsingh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u *User) NewEmail() {
	//This passes on a pointer/refernce object
	// func (u User) NewEmail() {
	// This passes on a copy object
	u.Email = "jitsingh1@redhat.com"
	fmt.Println("Email is:", u.Email)
}
