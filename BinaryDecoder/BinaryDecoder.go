// Chandler Dees
// CSC 4423
// 9 - 17 - 25

package main

import (
	"fmt"
	"strconv"
	//"encoding/binary"
)

func main(){
	// have to make user input string because int thats too large causes issues
	var input string // declaring user input 
	fmt.Print("Enter a binary string: ") 
	fmt.Scan(&input) // taking user input 

	fmt.Printf("Your input was: %s", input) // printing user input back at them
	fmt.Println() 

	// establishing default group size 
	grouping := 8 
	if len(input) % 7 == 0 {
		grouping = 7 // changing grouping to 7 if dealing with 7 bit ascii
	}

	// check if user input is a multiple of 8 or 7 and if not then pad as needed
	// if user input is a multiple of 8 or 7 proceed as normal
	if len(input) % grouping != 0 {
		padAmount := grouping - (len(input) %  grouping) // padding by the differnece between group amount and the input length
		for i := 0; i < padAmount; i++{
			// pad by adding zeros at the end 
			input = input + "0"
		}
	}

	// split user input into groups of 8 or 7 and then feed those in as bytes? to be interpreted into ASCII

	
	
	// grouping the total input into the grouping size  
	for i := 0; i < len(input); i += grouping {
		inputGroup := input[i : i + grouping]
		
		// converting input into integer then that is feed into ascii to become a character 

		val, err := strconv.ParseInt(inputGroup, 2, 64) // interpreting user input as binary then translating its value to dec
		// if the input cannot be interpreted then throw an error 
		if err != nil{
			fmt.Println("Error: ", err) 
			return 
		}

		// concatenating those characters to form a final output string 
		// printing the groups by formatting as chars and not printing a newline
		fmt.Printf("%c", val)

	}
	fmt.Println()
}

