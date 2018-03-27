//David Clarke G00335563
//Adapted from https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e

package main
//import
import (
	"fmt"
)

func intopost(infix string) string {
	//Creates a map with special characters, these are mapped to ints
	specials := map[rune]int{'*': 10, '.': 9, '|': 8}

	//Initializes an array of empty runes
	pofix := []rune{} 

	//Stack which stores operators from the infix regular expression
	s := []rune{}

	return string(pofix)
}

func main() {
	//Answer: ab.c*.
	fmt.Println("Infix:  ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer: abd|.*
	fmt.Println("Infix:  ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//Answer: abd|.c*.
	fmt.Println("Infix:  ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	//Answer: abb.+.c.
	fmt.Println("Infix:  ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))
} //main