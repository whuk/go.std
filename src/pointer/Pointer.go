package main
import "fmt"

func main() {

	var numPtr *int = new(int)
	fmt.Println(numPtr)

	*numPtr = 1
	fmt.Println(*numPtr)

	var num int = 1
	var numPtr2 *int = &num	// num의 주소값을 포인터 변수에 대입

	fmt.Println(numPtr2)
	fmt.Println(&num)

//	var numPtr3 *int = new(int)
//
//	numPtr3++ // 컴파일에러 포인터 연산은 허용하지 않음
//	numPtr3 = 0xc0820062d0	//  컴파일에러 주소도 직접대입안됨
//
//	fmt.Println(numPtr3)


}
