package main

import "fmt"

//struct groups related fiels into one type

type User struct{
	ID 		int
	Name 	string
	Email 	string
	Age 	int
}

func main(){

	u1 := User{
		ID: 1,
		Name: "Sangam",
		Email: "sangam@gmail.com",
		Age: 100,
	}
	
	fmt.Println(u1);

	
}