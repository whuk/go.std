package main
import "fmt"

func main() {

	strBttl := "bottle"
	strBttls := strBttl + "s"
	strOfbr := "of beer"
	strOnthWll := "on the wall"
	strOnWll := strOfbr + strOnthWll
	strTkdn := "Take one down, pass it around,"
	strNomore := "No more " + strBttls + strOnWll
	strLst := "Go to the store and buy some more"

	for i := 99; i > 0; i-- {
		switch  {
		case i > 1 && i <= 99:
			fmt.Println(i, strBttls, strOnthWll+ ",", i, strBttls, strOfbr)
			fmt.Println(strTkdn+ ",", i - 1, strBttls, strOnthWll+".")
		case i == 1:
			fmt.Println(i, strBttl, strOfbr, strOnthWll, ",", strNomore)
			fmt.Println(strNomore+ ",", strNomore)
			fmt.Println(strLst+ ",", 99, strBttls, strOnthWll+".")

		}
	}
}
