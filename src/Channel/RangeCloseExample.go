package main
import "fmt"

func main() {

	c := make(chan int)	// int 형 채널 생성

	go func(){
		for i := 0; i < 5; i++ {
			c <- i //채널에 값을 보냄
		}
		close(c)	// 채널을 닫음
	}()

	for i:= range c {	// range를 사용하여 채널에서 값을 꺼냄
		fmt.Println(i)
	}
}
