package main
import (
	"strconv"
	"fmt"
)

type Person struct {
	name string
	age int
}

func formatToString(arg interface{}) string {
	switch arg.(type){
	case Person:
		p := arg.(Person)
		return p.name + " " + strconv.Itoa(p.age)
	case *Person:
		p := arg.(*Person)
		return p.name + " " + strconv.Itoa(p.age)
	default:
		return "Error"
	}
}

func main() {

	fmt.Println(formatToString(Person{"Maria", 20}))
	fmt.Println(formatToString(&Person{"John", 12}))

	var andrew = new (Person)
	andrew.name = "Andrew"
	andrew.age = 35

	fmt.Println(formatToString(andrew))

	// 인터페이스에 저장된 타입이 어떤것인지 검사하려면

	var t interface{}
	t = Person{"Maria", 20}

	if v,ok := t.(Person); ok{
		fmt.Println(v, ok)
	}
}
