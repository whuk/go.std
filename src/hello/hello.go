package main

import "fmt"
import (
	"math"
	_ "time"
)

func main() {

	fmt.Println("Hello, world!")

	age := 10
	name := "Maria"

	_ = age
	_ = name

	var num1 int = 32
	var num2 int = -15
	var num3 int = 0723     // 8진수 0으로 시작
	var num4 int = 0xfa2c75 // 16진수 0x로 사작

	_ = num1
	_ = num2
	_ = num3
	_ = num4

	// 고정소수점
	var num5 float32 = 0.1
	var num6 float32 = .35
	var num7 float32 = 132.73287

	_ = num5
	_ = num6
	_ = num7

	// 부동소수점
	var num8 float32 = 1e7
	var num9 float64 = .12345E+2
	var num10 float64 = 5.32521e-10

	_ = num8
	_ = num9
	_ = num10

	var a float64 = 10.0

	for i := 0; i < 10; i++ {
		a = a - 0.1
	}

	fmt.Println(a)        //  라운딩에러   반올림 오차
	fmt.Println(a == 9.0) // 실수는 ==로 비교하면 안됨

	// 머신 앱실론 비교
	const epsilon = 1e-14                   // Go의 머신 앱실론
	fmt.Println(math.Abs(a-9.0) <= epsilon) // 연산값과 비교값(기대값)의 차이를 구한뒤 머신앱실론과 비교 (작거나 같아야함), 음수가 나올수있으므로 항상 절대값 붙일것

	// 복소수
	var num11 complex64 = 1 + 2i
	var num12 complex64 = 4.2342 + 2.3527i

	var num13 complex64 = 1.e+3i
	var num14 complex64 = 7.27151e-10i

	var num15 complex128 = 1 + 2i
	var num16 complex128 = 5.32521e-10 + .12345E+2i

	var r1 float32 = real(num11)
	var i1 float32 = imag(num11)

	var r2 float64 = real(num15)
	var i2 float64 = imag(num15)

	_ = num12
	_ = num13
	_ = num14
	_ = num16
	_ = r1
	_ = i1
	_ = r2
	_ = i2

	var num17 complex64 = complex(1, 2)
	var num18 complex128 = complex(5.32521e-10, .12345E+2)

	_ = num17
	_ = num18

	var num19 byte = 10   // 10진수
	var num20 byte = 0x32 // 16진수
	var num21 byte = 'a'  // 문자

	_ = num19
	_ = num20
	_ = num21

	var r11 rune = '한'
	var r12 rune = '\ud55c'
	var r13 rune = '\U0000d55c'
	fmt.Println(r11, r12, r13)

	// 숫자연산
	var num22 uint8 =3
	var num23 uint8 = 2

	fmt.Println(num22 + num23)
	fmt.Println(num22 - num23)
	fmt.Println(num22 * num23)
	fmt.Println(num22 / num23)
	fmt.Println(num22 % num23)
	fmt.Println(num22 << num23)	// 1100 8 + 4
	fmt.Println(num22 >> num23)	// 0
	fmt.Println(^num22) // 00000011 -> 11111100

	// 자료형

	var num24 int = 3
	var num25 float32 = 2.2

	fmt.Println(float32(num24) + num25)
	fmt.Println(num24 + int(num25))

	var num26 int32 = 80000
	fmt.Println(num26)
	fmt.Println(int16(num26))
}
