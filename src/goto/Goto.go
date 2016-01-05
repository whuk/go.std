package main
import "fmt"

func main() {

	var a int = 1

	if a == 1 {
//		fmt.Println("Error 1")
//		return
		goto ERROR1

		// 이 밑으로는 실행되지 않음
		b :=1
		fmt.Println(b)
	}

	if a == 2 {
//		fmt.Println("Error 2")
//		return
		goto ERROR2
	}

	if a==3 {
//		fmt.Println("Error 3")
//		return
		goto ERROR1
	}

	fmt.Println(a)
	fmt.Println("Sucess")

	return

	ERROR1 :
		fmt.Println("Error 1")
		return

	ERROR2 :
		fmt.Println("Error 2")
		return
}
