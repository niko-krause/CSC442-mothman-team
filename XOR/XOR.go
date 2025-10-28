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
	"flag"
	"fmt"
	"io"
	"os"
)

const DEBUG = true

// runtime debug flag (can be overridden by -debug)
var debugEnabled = DEBUG

// assignment constants
const keyPath = "key"
const bufSize = 4096

func main() {
	// flags: -demo to run demo mode, -debug to toggle debug prints
	demo := flag.Bool("demo", false, "run demo mode (prints binary representation of last input line)")
	debug := flag.Bool("debug", DEBUG, "enable debug prints")
	flag.Parse()

	debugEnabled = *debug

	// Default behavior: run assignment XOR stream mode (no flags)
	if *demo {
		// execute demo functions
		message := getMessage()
		toBinary(message)
		return
	}

	// assignment mode (default): read key from `key`, stream stdin/out
	xor()
}

func getMessage() string { // returns two things
	// taking in message through redirection (stole this from my program5)
	scanner := bufio.NewScanner(os.Stdin)

	var message string

	// getting stdin
	for scanner.Scan() {
		line := scanner.Text()
		message = line

		if debugEnabled {
			fmt.Printf("Message was: %s\n", message)
			fmt.Println()
		}
	}

	return message

}

func getKey() {
	// taking in key

	// convert key to binary with sprintf

	// return message and key
}

func toBinary(originalMessage string) string {
	binaryRep := ""
	// iterating through each character of message => binary
	for _, character := range originalMessage {
		binaryRep = fmt.Sprintf("%s%.8b", binaryRep, character)
	}

	if debugEnabled {
		fmt.Println("Binary Representation: ", binaryRep)
	}

	return binaryRep
}

func xor() { // assignment XOR: reads key from `key`, streams stdin, writes stdout
	// Read key file (binary)
	key, err := os.ReadFile(keyPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to read key file %q: %v\n", keyPath, err)
		os.Exit(1)
	}
	if len(key) == 0 {
		fmt.Fprintln(os.Stderr, "Error: key file is empty")
		os.Exit(1)
	}

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	buf := make([]byte, bufSize)
	keyLen := len(key)
	keyIndex := 0

	for {
		n, err := in.Read(buf)
		if n > 0 {
			for i := 0; i < n; i++ {
				buf[i] ^= key[(keyIndex+i)%keyLen]
			}
			keyIndex = (keyIndex + n) % keyLen

			if _, werr := out.Write(buf[:n]); werr != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to write output: %v\n", werr)
				os.Exit(1)
			}
		}

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "Error: failed to read stdin: %v\n", err)
			os.Exit(1)
		}
	}
}
