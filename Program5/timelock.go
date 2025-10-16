
/*
Author: Chandler Dees
Date: 10 - 16 - 25 
Description: Go program which implements the timelock algorithm
*/

package main 

// standard imports (stolen from my program 4)
import (
	"bufio"
	"fmt"
	//"io"
	//"net"
	"os"
	//"strconv"
	//"strings"
	//"time"
)

const DEBUG = true // toggles print statements and ability to input sys time


func main() {
	// need to collect user input 
	// takes this input from the cmd line at the same time as runtime 
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		
		// printing the input value 
		if DEBUG {
			line := scanner.Text()
			fmt.Printf("Read: %s\n", line)
		}
		
	}
		
	// collect system time 
	if DEBUG {
		// manually input desired system time 
	}
	
	
	
}
