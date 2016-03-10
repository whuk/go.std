package main
import "fmt"

func sum(a int, b int, c chan int){
	c <- a+b
}

func main() {

	c := make(chan int)		// int형 채널 생성

//	var d chan int
//	d = make(chan int)
//	fmt.Println(d)

	go sum(1, 2, c)	// sum을 고루틴으로 실행한뒤 채널을 매개변수로 넘김

	n := <-c

	fmt.Println(n)

}
