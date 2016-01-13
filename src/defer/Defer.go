package main
import "fmt"

//func hello() {
//	fmt.Println("Hello")
//}
//
//func world(){
//	fmt.Println("world")
//}

func helloWorld(){
	defer func(){
		fmt.Println("world")
	} ()

	func(){
		fmt.Println("Hello")
	} ()
}

func main() {
//
//	defer world()
//	hello()
//	hello()
//	hello()

	helloWorld()

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d", i)
	}
}
