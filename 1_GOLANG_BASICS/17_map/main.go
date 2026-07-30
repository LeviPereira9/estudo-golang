package main

import "fmt"

func main() {
	// map[keyType]valueType
	ages := map[string]int{
		"levi": 34,
		"john": 23,
	}

	fmt.Println(ages, len(ages));

	for k, v := range ages {
		fmt.Println(k, v);
	}

	//create empty map
	//make(map[K]V)

	var scores map[string]int;

	fmt.Println(scores, scores["a"]);

	scores = make(map[string]int);

	scores["levi"] = 1000;

	fmt.Println(scores)

	value, exist := scores["levi"];
	
	fmt.Println(value, exist)
	
	delete(scores, "levi");
	
	fmt.Println(scores)
	
	value, exist = scores["levi"];
	
	fmt.Println(value, exist)

}