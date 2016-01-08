package main
import (
	"fmt"
)

func main() {

//	var a [5]int
//	a[2] = 7
//	fmt.Println(a)

//	var a [5]int = [5]int{32,29,78,16,81}
//	var b = [5]int{32, 29, 78, 16, 81}
//	c := [5]int{32, 29, 78}
//
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)

//	a := [5]int{32, 24,78, 16, 4}
//	b := [5]int{
//		32,
//		5,
//		6,
//		7,
//		8,	// 여러줄일때는 마지막에 콬마를 붙임
//	}
//
//	c := [...]int{1, 2, 3, 4}
//	d := [...]string{"maria", "andrew", "jonh"}
//
//	fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Println(d)

	// 배열 순회
	a := [5]int{11, 12, 13, 14, 15}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 좀더 간단하게 레인지 사용하지
	for i, v := range a {
		fmt.Println(i, v)
	}

	// 두개를 써야 맞음
	for v := range a {
		fmt.Println(v)	// v로 해놨다고 value가 아니라 인덱스가 찍힘
	}

	for _, v := range a {
		fmt.Println(v)	// 이렇게 해야 value가 찍힘
	}

	// qodufqhrtk
	// 포인터나 주소값의 개념이 아니므로 그냥 :=.로 복사가능

	b := a
	fmt.Println(b)
}

