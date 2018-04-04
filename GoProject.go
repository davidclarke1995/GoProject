//David Clarke G00335563
//Adapted from https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
//Reference: https://golang.org/

//package main
package assets
 
func Intopost(infix string) string {
 	//Creates a map with special characters and maps them to ints
 	specials := map[rune]int{'*': 10, '.': 9, '|': 8}
 
 		case specials[r] > 0:
 			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
 				//Takes a character of the stack and puts it into pofix
 				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
 			}
 			//Less precendence then append current character onto the stack
 			s = append(s, r)
 		//Not a bracket or a special character
 		default:
 			//Takes the default characters and sticks it at the end of pofix
 			pofix = append(pofix, r)
 		}
 	}
 
 	//If there is anything left on the stack, append to pofix
 	for len(s) > 0 {
 		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
 	}
 	return string(pofix)
 } //intopost
 
 /*

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

}//main

*/

func TrimEndString(s string) string {
	if len(s) > 0 {
		s = s[:len(s)-2]
	}
	return s
}