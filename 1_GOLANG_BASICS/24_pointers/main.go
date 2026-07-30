package main

import "fmt"

func main() {

	// store the memory address of any value

	// &x -> address of x (makes a pointer)
	// *p -> deref (go to that address and read/write) [without returning it]

	score := 10
	score2 := score;
	fmt.Println("before: ", score)
	fmt.Println("before: ", score2)
	
	addScore(&score)
	
	fmt.Println("after: ", score)
	fmt.Println("after: ", score2)
}

func addScore(score *int) {
	*score += 5;
}
