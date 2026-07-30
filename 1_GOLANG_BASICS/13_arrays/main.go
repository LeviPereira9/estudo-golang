package main;

import (
	"fmt"
);

func main(){
	//fixed and can not grow
	 var marks [3]int;

	 marks[0] = 10;
	 marks[1] = 20;
	 marks[2] = 30;

	 // array literal
	 res := [5]int{2,3,4,5,6};

	 fmt.Println(len(res));

	 for index, value := range marks {
		fmt.Println(index, value);
	 }
}
