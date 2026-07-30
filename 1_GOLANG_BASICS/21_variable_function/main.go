package main

import "fmt"

func sumAll(nums ...int) int {

	sum := 0

	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {
	fmt.Println(sumAll(5, 5, 5, 5))

	values := []int{10, 20};

	fmt.Println(sumAll(values...));


	res := func(n int) int{
		return n * 2;
	}

	fmt.Println(res(2));

	// IIFE
	res1 := func(a int, b int) int {
		return a + b;
	}(1,2);

	fmt.Println(res1);
}
