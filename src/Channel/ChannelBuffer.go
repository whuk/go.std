package main
import (
	"runtime"
	"fmt"
)

func main() {

	runtime.GOMAXPROCS(1)

	done := make(chan bool, 2)	// 2개 채널 생성 -> 1개 이상이면 어싱크로 생성됨
	count := 4

	go func(){
		for i := 0; i < count; i++ {
			done <- true	// 채널에 true를 보냄
			fmt.Println("고루틴:", i)
		}
	}()

	for i := 0; i < count; i++ {
		<-done
		fmt.Println("메인함수:", i)
	}


}
