package main
import "fmt"

//func hello(n int){
//	n = 2 // 매겨변수에 2
//}
//
//func main() {
//
//	var n int = 1
//
//	hello(n)
//
//	fmt.Println(1)
//}


func hello(n *int){
	*n = 2
}

func main(){
	var n int = 1

	hello(&n)	// 주소값을 넘김

	fmt.Println(n)
}