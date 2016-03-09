package main
import (
	"fmt"
	"math/rand"
	"time"
)

//func hello(){
//	fmt.Println("Hello World!")
//}

func main() {

//	go hello()	// 함수를 고루틴으로 실행
//
//	fmt.Scanln()	// 메인함수가 종료되지 않도록 대기

	for i := 0; i < 100; i++ {
		go hello(i)	// 100 번 반복하며 고루틴 100개 생성
	}

	fmt.Scanln()
}


func hello(n int) {
	r := rand.Intn(100)	// 랜덤숫자 생성
	time.Sleep(time.Duration(r))	// 랜덤한 시간동안 대기
	fmt.Println(n)	// n 출력
}