/*
Author: Chandler Dees 
Date: 10 - 7 - 25 
Description: Connects to a server, extracts & prints overt and covert message
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const cutoff = 60 // milliseconds

func main() {
	// connect to the TCP server
	conn, err := net.Dial("tcp", "10.7.7.67:33333") // INSERT THE FOLLOWING: ipaddress:port
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}

	defer conn.Close()
	
	// opening message 
	fmt.Println("[connect to the chat server]")
	fmt.Println()

	reader := bufio.NewReader(conn)

	// keeps track of duration times
	var times []time.Duration

	// variable to hold the overt message
	var overtBuilder strings.Builder

	// keeps track of the delay
	var prevTime time.Time

	for {
		// getting the message
		start := time.Now()
		r, _, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				fmt.Println("\n[disconnect from the chat server]") // EDITED TO LINE UP WITH ASSIGNMENT OUTPUT 
				break
			}
			fmt.Println("\nError reading from server:", err)
			break
		}

		// record delay between characters
		if !prevTime.IsZero() {
			times = append(times, start.Sub(prevTime))
		}
		prevTime = start

		overtBuilder.WriteRune(r)
		fmt.Printf("%c", r)

		// stop once the end of the file has been reached
		if strings.HasSuffix(overtBuilder.String(), "EOF") {
			break
		}
	}


	// COMMENTING THIS OUT BECAUSE IT IS ABSENT FROM THE EXAMPLE OUTPUT ON THE ASSIGNMENT 
	/*
	// printing out the timing data
	fmt.Println("\nTiming data:")
	for _, t := range times {
		fmt.Printf("%.3f ", float64(t.Microseconds())/1000.0)
	}
	fmt.Println("\n")
	*/

	// convert timings to bits
	var bits strings.Builder
	for _, t := range times {
		if t.Milliseconds() > cutoff {
			bits.WriteByte('1') // 100ms-ish => 1
		} else {
			bits.WriteByte('0') // 25ms-ish => 0
		}
	}

	// ALSO COMMENTING THIS LINE OUT TO MAKE OUTPUT CLOSER TO EXAMPLE OUTPUT ON ASSIGNMENT 
	// printing bits for debugging
	//fmt.Println("Bits:", bits.String())

	// covert message is the ASCII rep of those bits
	covertMessage := bitsToString(bits.String())
	if strings.Contains(covertMessage, "EOF") {
		covertMessage = strings.Split(covertMessage, "EOF")[0]
	}

	fmt.Println("\nCovert message: ", covertMessage) // EDITED TO LINE UP WITH ASSIGNMENT OUTPUT 
}

// bitsToString tries all 8 bit alignments and returns the most readable one
func bitsToString(bits string) string {
	best := ""
	bestScore := -1

	for offset := 0; offset < 8; offset++ {
		var result strings.Builder
		score := 0 // readable characters

		// reading all bits at each offset
		for i := offset; i+8 <= len(bits); i += 8 {
			byteVal, _ := strconv.ParseInt(bits[i:i+8], 2, 64)
			b := byte(byteVal)
			// if the result falls into ASCII territory then it counts as readable 
			if b >= 32 && b <= 126 {
				result.WriteByte(b)
				score++
			} else {
				result.WriteByte('?')  // no score gained 
			}
		}
		
		// alignment with most printable characters is the winner
		if score > bestScore {
			best = result.String()
			bestScore = score
		}
	}
	return best
}
