//David Clarke G00335563
//Adapted from https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
//Reference: https://golang.org/

package main

import (
	"bufio"
	"fmt"
	"os"

	
)//imports

//test

func main() {

	option := 0
	exit := true

	for exit {//looping the user interface


		fmt.Print("Please select an option:\n1)Infix Expression to Postfix Expresssion\n2)Postix Regular Expresssion to NFA\n3)Exit\n")
		fmt.Scanln(&option)

		if option == 1 {//if statement taking in the users expression of choice for 1st option
			fmt.Print("Please enter the expression you want to convert: ")
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')
			expression = assets.TrimEndString(expression)

			//a.b.c* -> ab.c*.	(a.(b|d))* -> abd|.*	a.(b|d).c* -> abd|.c*.		a.(b.b)+.c -> abb.+.c.
			fmt.Println("Infix:  ", expression)
			fmt.Println("Postfix: ", assets.Intopost(expression))

		} else if option == 2 {//if statement taking in the users expression of choice for 2nd option

			fmt.Print("Please enter the expression to be converted:")
			reader := bufio.NewReader(os.Stdin)
			expression, _ := reader.ReadString('\n')

			expression = assets.TrimEndString(expression)

			fmt.Println("Postfix:  ", expression)
			fmt.Println("NFA: ", assets.Poregtonfa(expression))
			fmt.Print("\n")

			fmt.Print("Please enter a regular string to see if it matches the NFA:")
			//reader := bufio.NewReader(os.Stdin)
			regString, _ := reader.ReadString('\n')
			regString = assets.TrimEndString(regString)
			regString = assets.Intopost(regString) // Remove ending of string

			//"ab.c*|", "ccc"
			if assets.Pomatch(expression, regString) == true {
				fmt.Println("Regular string, ", regString, " matches the expression: ", expression)

			} else if assets.Pomatch(expression, regString) == false {
				fmt.Print("String does not match")
			} else {
				fmt.Print("An error has occured")
			}

		} else if option == 3 {
			fmt.Print("\nProgram Exited\n")
			exit = false
		} else {
			fmt.Print("\nPlease enter a valid option.\n")
		}
	}
} //main