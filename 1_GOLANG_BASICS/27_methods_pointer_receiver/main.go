package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{Name: "sangam", Age: 31}
	fmt.Println(u.Age);
	
	u.Birthday()

	fmt.Println(u.Age);
}

func (u *User) Birthday() {
	u.Age++
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}
