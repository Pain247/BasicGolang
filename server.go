//Server response "Hello world" bằng Go
package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I am D")
}
func f(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
func main() {
	go f(9)
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}
