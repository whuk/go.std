package main
import "fmt"


type Rectangle struct {
	//	width int
	//	height int

	width, height int
}

func NewRectangle(width, height int) *Rectangle {
	return &Rectangle{width, height}
}

func main() {

//	rect := NewRectangle(20, 10)
//	fmt.Println(rect)

	rect := &Rectangle{20, 10}
	fmt.Println(rect)
}
