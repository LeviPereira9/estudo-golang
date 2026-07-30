package main

import(
	"fmt"
	"strings"
)

func main(){
	firstName, lastName := "Levi", "Pereira da Silva"
	
	fullName := firstName + " " + lastName;

	fmt.Println(fullName);
	
	fmt.Println(strings.ToUpper(fullName))
	
}
