package main
import "fmt"

type Person struct {
	name string
	age int
}

func (p *Person	) greeting(){
	fmt.Println("Hello~")
}

type Student struct {
//	p Person	// 구조체를 가짐

	Person	// 임베딩

	school string
	grade int
}

// 오버라이딩 상황
func (p *Student) greeting(){
	fmt.Println("Hello Student~")
}

func main() {
	var s Student

//	s.p.greeting()
	s.Person.greeting()
	s.greeting()
}
