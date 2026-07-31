package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/hello", helloHandler)

	fmt.Println("try going to 8080 port")

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
	
}

func helloHandler(
	rw  http.ResponseWriter,
	re *http.Request){

	if re.Method != http.MethodGet {
		http.Error(rw, "Only get is allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = rw.Write([]byte("Hello from Go net/http server"))

}