//Tìm ước chung lớn nhất bằng GO
package main

import (
	"fmt"
)

func UCLN(x, y int) int {
	r := x % y
	if r == 0 {
		return y
	} else {
		return UCLN(y, r)
	}
}
func main() {
	k := 0
	k = UCLN(9, 3)
	fmt.Println(k)
}
