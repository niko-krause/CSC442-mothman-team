
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
	"time"
)

const DEBUG = true // toggles print statements and ability to input sys time
const layout = "2006 01 02 15 04 05" // arbitrary time to specify the format 


func main() {
	// need to collect user input 
	// takes this input from the cmd line at the same time as runtime 
	scanner := bufio.NewScanner(os.Stdin)
	
	var enoch string
	
	for scanner.Scan() {
		
		line := scanner.Text()	
		enoch = line
		
		if DEBUG {
			
			fmt.Printf("Read: %s\n", enoch)
		}
	}
	
	// parse the input enoch time as a time value 
	
	t, _ := time.Parse(layout, enoch)
	
	if DEBUG {
		fmt.Println("Parsed enoch time: ", t)
		fmt.Println()
	}
		
	// manually input desired system time 
	if DEBUG {
		
	}
	
	// collect system time 
	current := time.Now()
	fmt.Println("Current time is", current)
	
	// compare input with sys time to get ellapsed time
	
	elapsed := time.Since(t).Seconds()
	
	if DEBUG {
		fmt.Printf("Elapsed time is: %f seconds ", elapsed)
	}
	
	// do this with the comparisons built into the time package
	
	
	
	
	
	
}
