package main
import "fmt"

func sumr(a, b int) <- chan int {	// 리턴값이 받기전용 채널
	out := make(chan int)
	go func(){
		out <- a+b
	}()
	return out;
}

func main() {
	c := sumr(1,2)

	fmt.Println(<-c)

}
