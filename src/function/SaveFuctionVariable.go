package main
import "fmt"

func sumForSave(a int, b int) int {
	return a + b
}

func diffForSave(a int, b int) int {
	return a - b
}

func main() {

	var hello func(a int, b int) int = sumForSave
	world := sumForSave

	fmt.Println(hello(1,2))
	fmt.Println(world(3,4))

	 f:= []func(int, int) int {sumForSave, diffForSave}

	fmt.Println(f[0](1,2))
	fmt.Println(f[1](1,2))

	m := map[string]func(int, int) int {
		"sum" :sumForSave,
		"diff" : diffForSave,
	}

	fmt.Println(m["sum"](1,2))
	fmt.Println(m["diff"](2,2))
}
