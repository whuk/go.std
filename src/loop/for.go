package main
import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	Loop :
		// 여기에 뭐 들어오면 안됨...
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if j == 3 {
					break Loop
				}
				fmt.Println(i, j)
			}
		}
	fmt.Println("Hello world!")

	fmt.Println("------------------------------------------")

	for i, j := 0, 0; i < 10; i, j = i + 1, j + 2 {
		fmt.Println(i, j)
	}
}
