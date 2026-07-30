package main

import "fmt"

func main() {
	views := []int{10, 20, 45, 50, 60}

	for index, value := range views {
		fmt.Println(index, value);
	}

	sum := 0;

	for _, value := range views {
		fmt.Println(value)
		sum += value;
	}

	fmt.Println(sum, sum/len(views));
}
