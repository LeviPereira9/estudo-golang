package main;

import(
	"fmt"
);

func main(){

	for i := range 5 {
		fmt.Println(i)
	}
	/* 
	for v := range v {
		
	}

	for _, v := range v {
		fmt.Println(_);	}
	 */
	N := 5;
	sum := 0;

	for i := 0; i <= N; i++ {
		sum += i;
	}

	fmt.Println(sum);
}
