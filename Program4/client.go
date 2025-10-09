/* Author: Derrilon Young
* Date: 10/7/2025
* Description: Chat client that connects to a TCP server
* prints the overt message as it arrives, 
* and extracts a covert message based on timing delays 
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"time"
	"strings"
	"strconv"
)

const (
	cutoff = 60 //miliseconds 
)

func main(){
	//connect to the TCP server
	conn, err := net.Dial("tcp", "138.47.99.228:1337")
	if err != nil{
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	//prepare to close the connection 
	defer conn.Close()

	/*
	*	this client code allows a message to be typed out.
	*   that message will then be sent to the server which will read the message,
	*	and echo it back to the client. 
	*
	*/
	fmt.Println("Connected to server. Type a message and press Enter:")

	//read from stdin and send to server 
	reader := bufio.NewReader(conn)

	//keeps track of the duration times 
	var times []time.Duration

	//variable to hold the overt message 
	var overtMessage strings.Builder

	//keeps track of the delay 
	var prevTime time.Time 

	for {
		//getting the message
		start := time.Now()
		r, _, err := reader.ReadRune()

		if err != nil{
			if err == io.EOF{
				fmt.Println("\nServer closed Connection")
				break
			}
			fmt.Println("\nError reading from server:", err)
			break
		}

		if !prevTime.IsZero(){
			times = append(times, start.Sub(prevTime))
		}
		prevTime = start 

		overtMessage.WriteRune(r) //adds to the overt message 
		fmt.Printf("%c", r)

		//stop once the end of the file has been reached 
		if strings.HasSuffix(overtMessage.String(), "EOF"){//added
			break
		}

	}

	//prints out the final slice of time 
	fmt.Println("\nTiming data:")
	for _, t := range times {
		fmt.Printf("%d ", t.Milliseconds())
	}


	//decode the bits that were retrieved 
	var bits strings.Builder
	for _, t := range times{
		if t.Milliseconds() >= cutoff{
			bits.WriteByte('1')
		} else {
			bits.WriteByte('0')
		}
	}
	
	//convert the bits to ascii 
	covertMessage := bitsToString(bits.String())
	if strings.Contains(covertMessage, "EOF"){
		covertMessage = strings.Split(covertMessage, "EOF") [0]
	}

	fmt.Println("\nCovert message is:", covertMessage)
}

	func bitsToString(bits string) string {
		var results strings.Builder
		for i := 0; i+8 <= len(bits); i += 8{ 
			bytes := bits[i : i + 8]
			val, _ := strconv.ParseInt(bytes, 2, 8) 
			results.WriteByte(byte(val))
		} 
		return results.String()
	}
