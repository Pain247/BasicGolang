// Chương trình mới dừng ở mức tìm ngăn cách bởi dấu space chưa loại bỏ được các dấu . , !
package main

import (
	"fmt"
	"strings"
)

func count(arr []string, s string) int {
	k := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == s {
			k++
		}
	}
	return k
}
func WordCount(s string) map[string]int {
	var temp []string
	a := make(map[string]int)
	b := strings.Split(s, " ")
	for i := 0; i < len(b); i++ {
		if count(temp, b[i]) == 0 {
			temp = append(temp, b[i])
		}
	}

	for j := 0; j < len(temp); j++ {
		a[temp[j]] = count(b, temp[j])
	}
	return a
}

func main() {
	s := "Hello, I am D"
	fmt.Println(WordCount(s))
}
