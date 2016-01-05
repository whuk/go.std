package main
import "fmt"

func main() {

	a :=3
	b :=2
	c := a&b
	fmt.Printf("%08b", c)
	fmt.Println()

	d := 1
	e := &d

	fmt.Println(e)


	f := new(int)
	*f = 1
	fmt.Println(*f)


	g := make(chan int)

	go func() {
		g <- 1
	} ()

	h := <-g

	fmt.Println(h)
}
