//Tìm dãy Fibonacci không sử dụng đệ quy
package main

import "fmt"

func fibonacci(n int) []int {
	arr := make([]int, 0)
	a := 0
	b := 1
	arr = append(arr, a, b)
	temp := 0
	for b < n {
		temp = b
		b = b + a
		a = temp
		if b < n {
			arr = append(arr, b)
		}
	}
	return arr

}

func main() {
	arr := fibonacci(100000)
	fmt.Println(arr)
}
