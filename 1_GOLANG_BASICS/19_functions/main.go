package main

import "fmt"

func sumAndProduct(a int, b int) (int, int){
	sum := a + b;
	product := a * b;

	return sum, product;
}

func sum(nums ...int) int {
	sum := 0

	for _,n := range nums {
		sum += n;
	}

	return sum
}

func main() {
	fmt.Println(sum(1,2,3,4,5));
	fmt.Println(sumAndProduct(2,3));
}