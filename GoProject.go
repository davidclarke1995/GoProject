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
	
	//Loop over input
	for _, r := range infix {
		switch {

		case r == '(':
		//has open bracket at the end of the stack
			s = append(s, r)
		case r == ')':
			for s[len(s)-1] != '(' {
				pofix = append(pofix, s[len(s)-1]) //last element of 's'
				s = s[:len(s)-1]                   //everything in 's' except for the last character
			}

			s = s[:len(s)-1]
		//if a special character is found
		case specials[r] > 0:
			//for while the stack still has an element on it and the precedence of the current character that reads is <= the precedence of the top element of the stack
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				//Takes stack character and puts into pofix
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}

			s = append(s, r)
		//Not a bracket or a special character
		default:
			//Takes the default characters and puts it at the end of pofix
			pofix = append(pofix, r)
		}
	}

	//append to pofix whatever is left on the stack
	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

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