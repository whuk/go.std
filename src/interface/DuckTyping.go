package main
import (
	"fmt"
	"strconv"
)

type Duck struct {
	// 오리 구조체 정의
}

func (d Duck) quack(){
	fmt.Println("꽥~!")	// 덕의 울음 정의
}

func(d Duck) feathers(){
	fmt.Println("오리는 흰색과 회색털을 가지고 있다.")	// 생김새 정의
}

type Person struct {
	// 사람 구조체 정의
}

func(p Person) quack(){
	fmt.Println("사람은 오리흉내를 낸다 꽥~!")	// 사람도 콱을 정의
}

func (p Person) feathers() {
		fmt.Println("사람은 땅에서 깃털을 주워서 보여줍니다.")	// 사람도 reathers 정의
}

type Quacker interface {
	quack()
	feathers()	// 우는것 정의
}

func inTheForest(q Quacker){
	q.quack()
	q.feathers()
}

func f1(arg interface{})  {
	// 모든타입을 저장할수 있음
}

type Any interface{}	// 인터페이스에 메서드를 지정하지 않음

func f2(arg Any) {
	// 모든타입을 저장할수 있음
}

func formatString(arg interface{}) string {
	switch arg.(type) {

	case int:
		i :=arg.(int)
		return strconv.Itoa(i)

	case float32 :
		f := arg.(float32)
		return strconv.FormatFloat(float64(f), 'f', -1, 32)
	case float64 :
		f := arg.(float64)
		return strconv.FormatFloat(f, 'f', -1, 64)

	case string :
		s := arg.(string)
		return s

	default:
		return "Error"
	}
}

func main() {

	var donald Duck	// 오리
	var john	Person	// 사람

	inTheForest(donald)
	inTheForest(john)

	// 타입이 특정 인터페이스를 구현하고 잇는지를 검사 interface{}(인스턴스).(인터페이스)
	if v, ok := interface{}(donald).(Quacker); ok{
		fmt.Println(v, ok)
	}

	fmt.Println(formatString(1))
	fmt.Println(formatString(2.5))
	fmt.Println(formatString("Hello World!"))
}
