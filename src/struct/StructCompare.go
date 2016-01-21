package main
import "fmt"

type Rectangle struct {
	weight, height int
}

func rectangleScaleA(rect *Rectangle, factor int) {
	rect.weight = rect.weight * factor
	rect.height = rect.height * factor
}

func rectangleScaleB(rect Rectangle, factor int) {
	rect.weight = rect.weight * factor
	rect.height = rect.height * factor
}

func main() {

	rect1 := Rectangle{30, 30}
	rectangleScaleA(&rect1, 10)
	fmt.Println(rect1)

	rect2 := Rectangle{30, 30}
	rectangleScaleB(rect2, 10)
	fmt.Println(rect2)
}
