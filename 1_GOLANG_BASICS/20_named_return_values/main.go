package main

import "fmt"

func divide(a int, b int) (john int, sangam int) {
	john = a / b
	sangam = a + b

	return
}

func main() {
	j, s := divide(1, 2)

	fmt.Println(j, s);
}
