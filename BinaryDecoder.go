
// Chandler Dees 
// CSC 4423 
// 9 - 17 - 25 

package main 

import (
	"fmt" 
	//"encoding/binary"

)



func main(){
	var input string // declaring user input 
	fmt.Print("Enter a binary string: ") 
	fmt.Scan(&input) // taking user input 

	fmt.Printf("Your input was: %s", input) // printing user input back at them
	fmt.Println() 

	// check if user input is a multiple of 8 and if not then pad as needed
	// if user input is a multiple of 8 proceed as normal
	if !(len(input) % 8 == 0 || len(input) % 7 == 0) {
		print("Input length invalid, cannot decode to ASCII")
		return // exit
	}

	// split user input into groups of 8 or 7 and then feed those in as bytes? to be interpreted into ASCII

	// establishing default group size 
	grouping := 8 
	if len(input) % 7 == 0 {
		grouping = 7 // changing grouping to 7 if dealing with 7 bit ascii
	}
	
	// grouping the total input into the grouping size  
	for i := 0; i < len(input); i += grouping {
		inputGroup := input[i : i + grouping]
	}


	// converting input into integer then that is feed into ascii to become a character 


	// concatenating those characters to form a final output string 

	/*
	breaking the input into groups of 8 then turn them into bytes and feed those bytes into ascii for translation 
	*/

	
}

