package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://fstats.onrender.com/swagger-ui/index.html#/"

	resp, err := http.Get(url)

	if err != nil{
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("status code", resp.StatusCode)

	fmt.Println("status", resp.Status)
}
