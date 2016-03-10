package main
import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)
	count := 3

	go func(){
		for i := 0; i < count; i++ {
			done <- true	// 고루틴에 true를 보냄
			fmt.Println("고루틴 : ", i)
			time.Sleep(1*time.Second)	// 1초대기
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인 함수 : ", i)
	}
}
