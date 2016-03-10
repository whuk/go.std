package main
import "fmt"

func main() {

	c := make(chan int)

	go func(){
		c <- 1
	}()

//	close(c)	// 리턴값이두개, 채널이 닫힌경우와 열린경우 결과가 다름
	a, ok  := <-c
	fmt.Println(a, ok)


}
