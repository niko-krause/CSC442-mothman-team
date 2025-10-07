/* Author: Derrilon Young
* Date: 10/7/2025
* Description:
*
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"io"
	"time"
)

func main(){
	//connect to the TCP server
	conn, err := net.Dial("tcp", "localhost:1337")
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
		stop := time.Since(start)
		times = append(times, stop)
		fmt.Printf("%c", r)

	//printing out the final slice allows you to observe the durations
	fmt.Println("\nTiming data:")
	fmt.Println(times)

	}

	


}
