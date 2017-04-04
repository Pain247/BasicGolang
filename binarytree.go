package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

//Để sử dụng hàm String() trong package tree nên khai báo và sử dụng struct Tree của package tree thay vì sử dụng Tree tự tạo
//Nếu muốn sử dụng thì chỉ cần định nghĩa lại hàm String để trả về cây nhị phân dưới dạng string là ok :P
func Same(t1, t2 *tree.Tree) bool {
	fmt.Println(t1.String())
	fmt.Println(t2.String())
	if strings.Compare(t1.String(), t2.String()) == 0 {
		return true
	} else {
		return false
	}
}
func insert(t *tree.Tree, v int) *tree.Tree {
	if t == nil {
		return &tree.Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}
func main() {
	ch := make(chan int, 10)
	go Walk(tree.New(1), ch)
	go func() {
		k := len(ch)
		for i := 0; i < k; i++ {
			fmt.Print(" ", <-ch)
		}
	}()
	var t1, t2 *tree.Tree
	for i := 0; i < 10; i++ {
		t1 = insert(t1, i)
		t2 = insert(t2, i)
	}
	var bl bool
	bl = Same(t1, t2)
	fmt.Println(bl)
	var c string
	fmt.Scanln(&c)
}
