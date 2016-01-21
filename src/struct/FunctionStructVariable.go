package main
import "fmt"

type Ractangle struct {
	width, height int
}

func rectangleArea(rect *Ractangle) int {
	return rect.height * rect.width
}

func main() {

	rect := Ractangle{30, 30}
	area := rectangleArea(&rect)
	fmt.Println(area)
}
