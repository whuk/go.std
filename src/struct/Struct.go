package main
import "fmt"

type Rectangle struct {
//	width int
//	height int

	width, height int
}

func main() {

	var rect Rectangle
	fmt.Println(rect)

	var rect1 *Rectangle	// 구조체 포인터 선언
	rect1 = new (Rectangle)	// 메모리할당

	fmt.Println(*rect1)

	rect2 := new(Rectangle)	// 선언과 동시에 할당
	fmt.Println(rect2)

	var rect3 Rectangle = Rectangle{10, 20}
	fmt.Println(rect3)

	rect4 := Rectangle{30, 40}
	fmt.Println(rect4)

	rect5 := Rectangle{width:50, height:50}
	fmt.Println(rect5)

	var rect6 Rectangle
	var rect7 *Rectangle = new(Rectangle)
	rect6.height = 20
	rect7.height = 22

	fmt.Println(rect6)
	fmt.Println(rect7)
}
