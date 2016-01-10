package main
import "fmt"

func main() {

	// 코드양 줄이기, 클로저, 지연호출, 고루틴에 사용

	func(){ // 함수에 이름이 없음
		fmt.Println("Hello World!")
	} ()

	func(s string){
		fmt.Println(s)
	} ("Hello world!!")

	r := func(a int, b int) int {
		return a + b
	}(1, 2)	// 바로호출해서 r에 저장

	fmt.Println(r)
}
