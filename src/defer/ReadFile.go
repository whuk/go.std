package main
import (
	"os"
	"fmt"
)

func readHello(){
	file, err := os.Open("./src/defer/hello.txt")
	defer file.Close() // 함수가 종료될때 무조건 홀출되므로 에러처리같은데 많이 사용함

	if err != nil {
		fmt.Println(err)
		return	// 여기서 리턴하면 파일클로즈 호출함
	}

	buf := make([]byte, 100)
	if _, err = file.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(buf))
}

func main() {

	readHello()
}
