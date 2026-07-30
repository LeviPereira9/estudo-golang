package main

import (
	"fmt"
	"strings"
)


func main()  {
	var city string;
	city = "London";

	var peopleName = "Levi"; // inferred to string

	// ":=" = declarar váriavel e tipo ao mesmo tempo
	atoms := 2;
	atoms = atoms + 1000;
	atoms += 10;

	likes, comments := 10, 20;
	
	fmt.Println("Nome: ",peopleName)	
	fmt.Println("Cidade", city)	
	fmt.Println("Atomos na praia = ", atoms)
	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("Curtidas: ", likes)
	fmt.Println("Comentários", comments)
}
