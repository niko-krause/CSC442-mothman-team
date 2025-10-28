
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
const layout = "2006 01 02 15 04 05" // using DateTime layout from time package 


func main() {
	timeCalculations()
}

/*
taking the users input through redirection and does the following:
parses the provided enoch time as a time value rather than a string 
takes the current system time 
compares the two to get the seconds passed 
returns second value
*/

func timeCalculations(){
	// need to collect user input 
	// takes this input from the cmd line at the same time as runtime 
	scanner := bufio.NewScanner(os.Stdin)
	
	var enoch string
	
	// getting stdin 
	for scanner.Scan() {
		
		line := scanner.Text()	
		enoch = line
		
		if DEBUG {
			
			fmt.Printf("Read: %s\n", enoch)
		}
	}
	
	// location information 
	
	location, _ := time.LoadLocation("America/Chicago")
	
	// parse the input enoch time as a time value 
	
	parsedEnoch, _ := time.ParseInLocation(layout, enoch, location) // removed .UTC
	//parsedEnoch = parsedEnoch.UTC()
	
	if DEBUG {
		fmt.Println("Parsed enoch time: ", parsedEnoch)
		fmt.Println()
	}
		
	
	
	// collect system time 
	current := time.Now() // removed .UTC()
	
	/*
	// manually input desired system time 
	if DEBUG {
		currentStr := "2017 10 01 00 00 00" 
		// need to parse this bad boy also 
		current, _ = time.Parse(layout, currentStr)
	} 
	*/
	
	fmt.Println("Current time is", current)
	
	// compare input with sys time to get ellapsed time
	
	//elapsed := time.Since(parsedEnoch).Seconds() 
	
	elapsed := current.Sub(parsedEnoch).Seconds()
	
	if DEBUG {
		fmt.Printf("Elapsed time is: %v seconds \n", elapsed)
	}
	
}
