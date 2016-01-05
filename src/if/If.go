package main
import (

	"fmt"
	"io/ioutil"
)

func main() {

	var b []byte
	var err error
	b, err = ioutil.ReadFile("./src/if/hello.txt")


	if err == nil {
		fmt.Printf("%s", b)
	}else {
		fmt.Println(err)
	}

	if b,err := ioutil.ReadFile("d:/goworkspace/go.std/src/if/hello.txt"); err == nil {
		fmt.Printf("%s", string(b))
	} else {
		fmt.Println(err)
	}


}
