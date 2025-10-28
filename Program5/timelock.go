
/*
Author: Chandler Dees
Date: 10 - 16 - 25 
Description: Go program which implements the timelock algorithm
*/

package main 

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"crypto/md5"
	"io"
)

const DEBUG = true // toggles print statements and ability to input sys time
const layout = "2006 01 02 15 04 05" // using DateTime layout from time package 


func main() {
	finalElapsed := timeCalculations()
	doubleHash(finalElapsed)
	
}


/*
taking the users input through redirection and does the following:
parses the provided enoch time as a time value rather than a string 
takes the current system time 
compares the two to get the seconds passed 
returns second value
*/


func timeCalculations() float64{ // removed time.Time
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
	
	// location information (needed to fix daylight savings issues)
	
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
	

	// manually input desired system time 
	if DEBUG {
		currentStr := "2017 10 01 00 00 00" 
		// need to parse this bad boy also 
		current, _ = time.Parse(layout, currentStr)
	} 

	
	fmt.Println("Current time is", current)
	
	// compare input with sys time to get ellapsed time
	
	elapsed := current.Sub(parsedEnoch).Seconds()
	
	if DEBUG {
		fmt.Printf("Elapsed time is: %v seconds \n", elapsed)
	}
	
	return elapsed
	
}

/*
function that takes in the elapsed time from epoch and then 
computes the md5 hash twice, extracts the first two chars 
and backwards first two ints, concatenates, and returns 

*/

func doubleHash (elapsed float64){
	
	// need to convert the ellapsed seconds into string
	elapsedString := fmt.Sprintf("%f", elapsed)
	
	// first hash
	hash := md5.New()
	
	io.WriteString(hash, elapsedString)
	
	firstHash := hash.Sum(nil)
	
	// second hash
	hashTwo := md5.New() 
	
	io.WriteString(hashTwo, fmt.Sprintf("%x", firstHash))
	
	secondHash := hashTwo.Sum(nil)
	
	fmt.Printf("%x", secondHash)
	
}
