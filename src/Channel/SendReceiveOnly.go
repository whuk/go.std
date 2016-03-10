package main
import "fmt"

// 전용채널에서 다른 행동을 하면 컴파일 에러 발생
func producer(c chan <-int) {
	// 보내기 전용 채널

	for i := 0; i < 5; i++ {

			c <- i

	}
	c<-100
}

func consumer(c  <- chan int) {
	// 받기전용 채널
	for i:=range c{
		fmt.Println(i)
	}

	fmt.Println(<-c)
}

func main() {
	c := make(chan int)



	go producer(c)
	go consumer(c)

	fmt.Scanln()
}
