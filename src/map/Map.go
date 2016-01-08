package main
import "fmt"

func main() {

//	var a map[string]int	//  키가 스트링, 벨류가 인트
//	var a map[string]int = make(map[string]int)
//	var b = make(map[string]int)
//	c := make(map[string]int)

//	a := map[string]int{"Hello" :10, "world": 20}
//
//	b := map[string]int{
//		"Hello":10,
//		"world":20,	// 역시 세로는 마지막에 컴마붙임
//	}
//	fmt.Println(a)
//	fmt.Println(b)

	// 맵에 데이터 저장하고 조회하기
//	solarSystem := make(map[string]float32)
//
//	solarSystem["Mercury"] = 87.969
//	solarSystem["Venus"] = 224.70069
//	solarSystem["Earth"] = 365.25641
//	solarSystem["Mars"] = 686.9600
//
//	fmt.Println(solarSystem["Earth"])
//	fmt.Println(solarSystem["Pluto"])	// 없는키를 조회하면 빈값잋 ㅜㄹ력됨
//
//	value, ok := solarSystem["Piuto"]
//	fmt.Println(value, ok)
//
//	if value, ok := solarSystem["Earth"]; ok {
//		fmt.Println(value)
//	}
//
//	// 맵의 순회
//	for key, value := range solarSystem {
//		fmt.Println(key, value)
//	}
//
//	for _, value := range solarSystem {
//		fmt.Println(value)
//	}
//
//	for key, _ := range solarSystem {
//		fmt.Println(key)
//	}

	// 맵에서 데이터 삭제
	a := map[string]int{"Hello":10, "world":20}
	delete(a, "world")
	fmt.Println(a)

	// 맵안에 맵 만들기
	terrestrialPlanet := map[string]map[string]float32{
		"Mercury":map[string]float32{
			"meanRadius":2439.7,
			"mass": 3.3022E+23,
			"orbitalPeriod":87.969,
		},
		"Venus":map[string]float32{
			"meanRadius":6051.8,
			"mass": 4.8676E+24,
			"orbitalPeriod":224.70069,
		},
	}

	fmt.Println(terrestrialPlanet["Mercury"]["mass"])

}
