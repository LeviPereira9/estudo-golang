package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	//go dont use exceptions for normal failures. :O
	// functions -> return errors as normal return values

	// val, err := somenthing();
	// if err != nil {handle the err}
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	input := "3a";

	level, err := parseLevel(input);

	if err != nil {
		return err;
	}

	fmt.Println("select level", level);

	return nil;
}

func parseLevel(s string) (int, error) {
	// (value, err)
	// nil error -> success
	// non ill -> failure

	// pattern
	n, err := strconv.Atoi(s);

	if err != nil {
		return 0, fmt.Errorf("level must be a number");
	}

	if n < 1 || n > 5 {
		return 0, fmt.Errorf("level must be between 1 and 5");
	}

	return n, nil;
}
