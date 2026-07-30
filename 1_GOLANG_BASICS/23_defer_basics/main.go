package main

import (
	"errors"
	"fmt"
)

func main() {

	// defer resp.body.Close();

	fmt.Println("Case 1: success")
	if err := doWork(true); err != nil {
		fmt.Println("error: ", err)
	}
	
	fmt.Println("Case 2: success")
	if err := doWork(false); err != nil {
		fmt.Println("error: ", err)
	}
}

func doWork(success bool) error {
	// resource related
	// start message -> resource acquired
	//cleanup message -> resource released

	fmt.Println("start: resource acquired")

	defer fmt.Println("cleanup: resource released");

	if !success{
		return errors.New("Something went wrong, returning early")
	}

	fmt.Println("workd: doing something imp")
	fmt.Println("work: this work is done")

	return nil
}
