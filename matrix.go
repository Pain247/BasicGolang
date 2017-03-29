//Chương trình tạo ma trận bằng slice 2 chiều
package main
import(
	"fmt"
)
type MT struct{
	r,c int
	arr [][]int  
}
//Test interface, giống với khai báo nguyên mẫu hàm trong C
type MTer interface{
	alloc()
	nhap()
	print() 
	add(mt1 *MT, mt2 *MT) MT
}


func main(){
	var A,B,C MT
	A.r = 3
	A.c = 3
	B.r = 3
	B.c = 3
	//Sử dụng A như một đối tượng tham chiếu đến phương thức bên trong interface thay vì truyền tham chiếu vào hàm alloc()
	A.alloc()
	A.nhap()
	B.alloc()
	B.nhap()
	C = add(&A,&B)
	A.print()
	B.print()
	C.print()
}
func (mt *MT) alloc(){
	mt.arr = make([][]int, mt.r)
	for i := 0; i < mt.r; i++ {
		mt.arr[i] = make([]int, mt.c)
	}
}
func (mt *MT) nhap(){
	 	for i := 0; i < mt.r; i++ {
		for j := 0; j < mt.c; j++ {
			fmt.Printf("Nhap vao phan tu [%d][%d] :",i,j)
			fmt.Scan(&mt.arr[i][j])
		}
	}
}
func (mt *MT) print(){
	for i:=0;i< mt.r ; i++{
		for j:=0;j < mt.c;j++{
			fmt.Printf("%d ", mt.arr[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}
func add(mt1 *MT, mt2 *MT) MT{
	var mt3 MT
	mt3.r = mt1.r
	mt3.c = mt1.c
    mt3.alloc()
	 for i := 0; i < mt3.r; i++ {
		for j := 0; j < mt3.c; j++ {
			mt3.arr[i][j] = mt1.arr[i][j] + mt2.arr[i][j]
		}
	}
	return mt3

}