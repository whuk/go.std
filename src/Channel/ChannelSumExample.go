package main
import "fmt"

func num(a, b int) <- chan int{
	out := make(chan int)

	go func(){
		out <- a
		out <- b
		close(out)
	}()

	return out
}

func sumc(c <-chan int) <- chan int {
	out := make(chan int)

	go func(){
		r :=0
		for i:=range c {
			r = r+i
		}
		out <- r
	}()
	return out
}

func main() {

	c := num(1,2)	// 1, 2 가들어있는 채널이 리턴됨
	out := sumc(c)	// 채널 안에있는값이 더해저서 그게 들어있는 채널 리턴
	fmt.Println(<-out)

}
