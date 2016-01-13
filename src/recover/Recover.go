package main
import "fmt"

func f() {
	defer func(){
		s := recover()	// 리커버는 반드시 지연호출로 사용해야함

		fmt.Println(s)
	} ()

	panic("Error!")
}

func main() {

	f()
	fmt.Println("Hello world")

//	 a:= [...]int{1,2}
//	for i := 0; i < 3; i++ {
//		fmt.Println(a[i])
//	}

//	panic("Error!!")
//	fmt.Println("Hello world")
}
