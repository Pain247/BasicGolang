package main

import (
	"fmt"
)
//Hàm có thể coi như là đối số truyền đc vào trong hàm khác
func calc(test func(float32, float32) float32) float32 {
	return test(3, 4)
}
func main() {
	k := func(x, y float32) float32 {
		return x * y
	}
	fmt.Println(k(2, 3))
}
