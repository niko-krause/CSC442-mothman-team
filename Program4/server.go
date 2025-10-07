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
	"strings"
	"time"
)

const (
	address = "" //allow all connections 
	port = "1337" 
	zero = 25 //delay miliseconds for 0 
	one = 100 //delay miliseconds for 1 
)

func main(){
	listener, err := net.Listen("tcp", address+":"+port)
	if err != nil{
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on %s:%s\n", address, port)

	for {
		conn, err := listener.Accept()
		if err != nil{
			fmt.Println("Error accepting connection:", err)
			continue
		}
		fmt.Println("New client connected!")
		go sendMessage("This is a really long message that will be sent a charater at a time", conn)

	}
}

func getCovert(){
	covert := "secret" + "EOF"
	var result strings.Builder
	for _, c := range covert {
		result.WriteString(fmt.Sprintf("%08b", c))
	}
	return result.String()
}

func sendMessage(message string, conn net.Conn){
	defer conn.Close()
	covert := getCovert()

	for i, ch := range message{
		_, err := conn.Write([]byte(string(ch)))
		if err != nil{
			fmt.Println("Error writing to client:", err)
			return
		}

		//transmit timing based on covert bit 
		if covert[i%len(covert)] == '1'{
			time.Sleep(one * time.Millisecond)
		} else {
			time.Sleep(zero * time.Millisecond)
		}
	}

	fmt.Println("Overt message sent. Closing connection.")
}