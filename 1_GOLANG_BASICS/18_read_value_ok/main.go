package main

import "fmt"

func main() {
	points := map[string]int{
		"a": 10,
		"b": 0,
	}

	fmt.Println("a", points["a"]);
	fmt.Println("b", points["b"]);
	fmt.Println("c", points["c"]);

	valB, okB := points["b"];

	fmt.Println(valB, okB);

	for _,v := range points {
		fmt.Println(v)
	}
}
