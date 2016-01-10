package main
import "fmt"


//func hello(){
//	fmt.Println("Hello world!")
//}

func main() {
	hello()
}

//func sum(a int, b int) int {
//	return a+b
//}

func sum(a int, b int) (r int) {
	// 리턴값 지정가능
	r = a + b
	return // 여기에다가는 지정하지 않음
}

// 여러개 사용하기
func sumAndDiff(a int, b int) (int, int){
	return a+b, a-b
}

func helloM() (int, int, int, int, int){
	return 1, 2, 3, 4, 5
}

func sumAndDiffM(a int, b int) (sum int, diff int){
	sum = a + b
	diff = a - b
	return
}

func sumMulti(n ...int) int {
	total := 0
	for _, v := range n  {
		total += v
	}

	return total
}

func hello(){
	// C와 달리 위치 상관없음
	fmt.Println("Hello world!")

	r := sum(1, 2)
	fmt.Println(r)

	_, diff := sumAndDiff(1, 2)
//	fmt.Println(sum)
	fmt.Println(diff)

	a, _, c, _, e := helloM()
	fmt.Println(a, c, e)

	sum, diff := sumAndDiffM(6, 2)

	fmt.Println(sum, diff)

	fmt.Println(sumMulti(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	// 가변인자는 슬라이스타입이므로 슬라이스를 바로 넘겨도 됨
	numbers := []int {1,2,3,4,5}
	fmt.Println(sumMulti(numbers...))	// 단, ...을 써줘야함..
}