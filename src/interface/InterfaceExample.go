package main
import "fmt"

type hello interface {

}

type MyInt int	// int형을 MyInt로 정의

func (i MyInt) Print(){
	fmt.Println(i)	// MyInt에 Print 메서드 연결
}



type Rectangle struct {
	width, height int
}

func (r Rectangle) Print(){
	fmt.Println(r.width, r.height)
}

type Printer interface {
	Print()
}

func main() {
	var h hello
	fmt.Println(h)

	var i MyInt = 5

	var p Printer // 인터페이스 선언

	p = i
	p.Print()

	r := Rectangle{10, 20}

	p = r
	p.Print()

	p1 := Printer(i)
	p1.Print()

	p2 := Printer(r)
	p2.Print()

	// 슬라이스로 선언
	pArr := []Printer{i, r}
	for index,_:= range pArr {
		pArr[index].Print()	// 인덱스순회
	}

	for _,value := range pArr  {
		value.Print()	// value 순회
	}

}
