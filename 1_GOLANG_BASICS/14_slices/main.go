package main

import "fmt"

func main() {
	// common colection type
	// dynamic and can grow
	// []type{...}

	results := []string{"levi", "abduh"};


	fmt.Println(results);

	results = append(results, "aham");
	fmt.Println(results);
	
}
