package main
import "fmt"

func main() {

	const (
		Sunday = 0
		Monday = 1
		Tuesday = 2
		Wednesday = 3
		Thursday = 4
		Friday = 5
		Saturday = 6
		numberOfDays = 7
	)

	const (
		Zero = iota
		one
		two
		three
		four
		five
	)

	fmt.Println(three)

	const (
		a = 1 << iota
		b = 1 << iota
		c = 1 << iota
	)

	fmt.Println(c)

	const (
		bit0, msk0 = 1 << iota, 1 << iota - 1
		bit1, msk1
		_,_
		bit3, msk3
	)

	fmt.Println(bit3, msk3)
}
