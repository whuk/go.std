package main
import (
	"runtime"
	"fmt"
)

func main() {
//	runtime.GOMAXPROCS(runtime.NumCPU())	// CPU개수구한뒤에 프로세스갯수설정
//	fmt.Println(runtime.GOMAXPROCS(0))	// 설정값 출력

	runtime.GOMAXPROCS(1)
	s := "Hello World!"

//	for i := 0; i < 100; i++ {
//		go func(n int) {
//			// 익명함수를 고루틴으로 실행	클로저
//			fmt.Println(s, n)	// 매개변수 n 출력
//		}(i)	// 반복문의 변수를 매개변수로 넘김
//	}

	for i := 0; i < 100; i++ {
		go func(){
			fmt.Println(s, i)	// 반복문의 변수를 클로저에서 바로 출력
		}()	// 이렇게하면 루프다돌고 찍히므로 i는 100밖에 안나옴(고루틴이 반복문이 끝나고 생성되기 때문)
	}

	fmt.Scanln()
}
