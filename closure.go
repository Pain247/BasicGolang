// Dãy Fibonacci sử dụng tính chất của Closure function
package main

import "fmt"
var (
	arr = []int{0, 1}
)

func fibonacci() func() int {
	return func() int {
		arr = append(arr, arr[len(arr)-2]+arr[len(arr)-1])
		return arr[len(arr)-2] + arr[len(arr)-1]

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
