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

	//create a slice of time.Duration objects so that the times can be grouped
	var times []time.Duration

	//variable to hold the overt message 
	var overt strings.Builder

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
		elapsed := time.Since(start)
		times = append(times, stop)
		overt.WriteRune(r) //writes the message to the overt variable
		fmt.Printf("%c", r)

		//stop once the end of the file has been reached 
		if strings.HasSuffix(overt.String(), "EOF"){//added
			break
		}

	}
	//printing out the final slice allows you to observe the durations
	fmt.Println("\nTiming data:")
	fmt.Println(times)


	//decode the bits that were retrieved 
	var bits strings.Builder 
	for _, t := range times{
		if t > threshold { //defind the threshold later
			bits.WriteByte('1')
		} else {
			bits.WriteByte('0')
		}
	}
	
	//convert the bits to ascii 


}
