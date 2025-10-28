
/*
  Author: Chandler Dees 
  Date: 10 - 28 - 25 
  Description: XOR Crypto: takes in message and key, XORs them together
			   one bit at a time with a buffer

*/

package main 

// test this program because it works both ways 
// if you xor it twice you w/ same key should get the initial back 
// so it should be able to decode and encode 
// google that and see his discord messages

import (
	"bufio"
	"fmt" 
	"os"

)

const DEBUG = true 

func main() {
	// execute other functions 
	message := getMessage()
	toBinary(message)
	
	
}

func getMessage() string {// returns two things 
	// taking in message through redirection (stole this from my program5)
	scanner := bufio.NewScanner(os.Stdin)
	
	var message string 
	
	// getting stdin 
	for scanner.Scan() {
		line := scanner.Text() 
		message = line 
		
		if DEBUG {
			fmt.Printf("Message was: %s\n", message)
			fmt.Println()
		}
	}
	
	return message 
	
} 

func getKey(){
	// taking in key 
	
	// convert key to binary with sprintf 
	
	// return message and key 
}

func toBinary(originalMessage string) string{
	binaryRep := "" 
	// iterating through each character of message => binary 
	for _, character := range originalMessage {
		binaryRep = fmt.Sprintf("%s%.8b", binaryRep, character)
	}
	
	if DEBUG {
		fmt.Println("Binary Representation: ", binaryRep)
	}
	
	
	return binaryRep
}


func xor(){ // takes in message and key as inputs values
	
	// xor together
	
	// return this output in a separate file 
}

