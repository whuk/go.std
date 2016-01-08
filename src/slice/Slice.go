package main
import "fmt"

func main() {

	// 슬라이스는 배열과 같지만, 동적, 리스트로 생각하면됨, 레퍼런스 타입임

//	var a []int	// 선언
//	fmt.Println(a)

	// make 함수를 써서 공간할당 해야함
//	var a []int = make([]int, 5)	// make 함수로 int 형 길이 5 할당
//	var b = make([]int, 5)	// 생략 가능
////	c := make([], 5)	// 더생략 --> 이거안됨
//
//	fmt.Println(a)
//	fmt.Println(b)
////	fmt.Println(c)

//	a := []int{1,2,3,4,5}
//	b := []int{
//		6,
//		7,
//		8,
//		9,
//		0,	// 역시 세로로할땐 마지막에도 , 붙임
//	}
//	fmt.Println(a)
//	fmt.Println(b)

	// 용량할당가능 어쨋든 가변길이
//	a := make([]int, 5, 10)
//	fmt.Println(len(a))
//	fmt.Println(cap(a))

	// 슬라이스에 값 추가하기
//	a := []int{1,2,3}
//	a = append(a, 4,5,6)
//	fmt.Println(a)

//	a := []int{1,2,3}
//	b := []int{4,5,6}
//	a = append(a, b...)
//	fmt.Println(a)

//	a := [3]int{1,2,3}
//	var b [3]int
//	b = a	// 배열복사
//	b[0] = 9	// 복사후 값변경
//	fmt.Println(a)
//	fmt.Println(b)

//	a := []int{1,2,3}
//	var b []int
//
//	b = a
//	b[0] = 9
//
//	fmt.Println(a)
//	fmt.Println(b)

	// 슬라이스의 복사는 copy를 사용한다
//	a := []int{1,2,3,4,5,6,7}
//	b := make([]int, 3)	// 메이크로 할당
//
//	copy(b, a)
//
//	fmt.Println(a)
//	fmt.Println(b)

//	a := []int{1,2,3}
//	b := make([]int, 3)
//
//	copy(b,	a)
//	b[0] = 9
//
//	fmt.Println(a)
//	fmt.Println(b)

	// 슬라이스와 용량
//	a := []int{1,2,3,4,5}
//	fmt.Println(len(a), cap(a))
//	a = append(a, 6,7,8)
//	fmt.Println(len(a), cap(a))

	// 부분슬라이스
//	a := []int{1,2,3,4,5}
//	b := a[0:5]
//	fmt.Println(a)
//	fmt.Println(b)

//	a := []int{1,2,3,4,5}
//	fmt.Println(a[:])
//	fmt.Println(a[0:])
//	fmt.Println(a[:5])
//	fmt.Println(a[0:len(a)])
//
//	fmt.Println(a[3:])
//	fmt.Println(a[:3])
//	fmt.Println(a[1:3])

	// 배열에서 부분가능하지만 이건 참조임
//	a := [5]int{1,2,3,4,5}
//	b := a[:2]
//	b[0] = 9
//	fmt.Println(a)
//	fmt.Println(b)

	// 슬라이스 용량지정도 가능
	a := []int{1,2,3,4,5,6,7,8}
	b := a[0:6:8]	// 마지막인자는 용량, 근데 기존슬라이스의 용량을 넘을수는 없음
	fmt.Println(len(b), cap(b))
	fmt.Println(b)
}
