package main

import "fmt"

func main() {
	//make([]T, len, cap)
	scores := make([]int, 0, 5)

	fmt.Println(scores, len(scores), cap(scores));

	todos := []string{"do uw", "ads"};

	more := []string{"yes", "no"};

	todos = append(todos, more...);

	fmt.Println(todos);

	// ... desempacota slices
	// ... em funções significa que os parametros são variaveis.
	//fmt.Println(more...)
}
