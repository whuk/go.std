package main
import "fmt"

func factirial(n uint64) uint64 {
	if n == 0 {
		return 1
	}

	return n * factirial(n-1)
}

func main() {

	fmt.Println(factirial(5))
}
