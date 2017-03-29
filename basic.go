//Test input/output
package main
import (
	"fmt"
	"bufio"
	"os"
)

func main(){
reader := bufio.NewReader(os.Stdin)
a,_ := reader.ReadString('\n')
fmt.Println(a)
fmt.Println("Nhap vao:")
text := 0
fmt.Scanln(&text) //Truyền tham chiếu khi dùng tham chiếu như trong C
fmt.Println(text)
fmt.Println(text)
}
