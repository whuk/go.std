package main
import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var s1 string = "한글"
	var s2 string = "Hello"

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Println(len(s1))
	fmt.Println(len(s2))

	fmt.Println("한글길이", utf8.RuneCountInString(s1))

	var s3 string ="한글"
	var s4 string ="한글"
	var s5 string ="Go"

	fmt.Println(s3 == s4)
	fmt.Println(s3+s4+s5)
	fmt.Println("한녕하세요"+s3)
}
