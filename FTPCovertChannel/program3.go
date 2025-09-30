/*
Authors: Chandler Dees
Date: 9/25/25
*/
package main

import (
	"fmt"
	"log"
	"os"

	//"program3.go/BinaryDecoder"
	"strconv"

	"github.com/secsy/goftp"
)

const (
	address = "138.47.99.229" // since we are logging in on where its hosted
	// otherwise we would put the IP address of where we are logging in
	port     = "21" // port number
	username = "percypatterson"
	password = "himalayas"
	path     = "/files/10" // find by pwd
	// where are the files you are interested in
)

// METHOD = false for 7-bit mode
// METHOD = true for 10-bit mode
var METHOD = true

// BinaryPerms converts file permissions into a binary string.
// If METHOD=false, only last 7 bits are kept, skipping entries where the first 3 bits are set.
// If METHOD=true, all 10 bits are included.
func BinaryPerms(entries []os.FileInfo) string {
	output := ""

	for _, entry := range entries {
		modeString := entry.Mode().String() // e.g. "-rw-r--r--"
		binary := ""

		// permission string to binary
		for _, bit := range modeString {
			if bit == '-' {
				binary += "0"
			} else {
				binary += "1"
			}
		}

		if METHOD {
			// use all 10 bits
			output += binary
		} else {
			// slippa
			if binary[0] == '1' || binary[1] == '1' || binary[2] == '1' {
				continue
			}
			// use only the last 7 bits
			last7 := binary[len(binary)-7:]
			output += last7
		}
	}

	return output
}

// convert input into 7 bit ascii
func Decode(inputString string) string {
	const grouping = 7
	decodedString := ""

	// pad to a multiple of 7
	if len(inputString)%grouping != 0 {
		padAmount := grouping - (len(inputString) % grouping)
		for i := 0; i < padAmount; i++ {
			inputString += "0"
		}
	}

	// break into groups of 7
	for i := 0; i < len(inputString); i += grouping {
		inputGroup := inputString[i : i+grouping]

		val, err := strconv.ParseInt(inputGroup, 2, 64)
		if err != nil {
			log.Println("Error decoding group:", inputGroup, err)
			continue
		}
		decodedString += fmt.Sprintf("%c", val)
	}

	return decodedString
}

func main() {
	// FTP client config
	config := goftp.Config{
		User:     username,
		Password: password,
	}

	client, err := goftp.DialConfig(config, address+":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// get dir listing
	entries, err := client.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// convery permissions to binary
	binver := BinaryPerms(entries)

	// decode binary into ASCII
	result := Decode(binver)

	// final result
	fmt.Println(result)
}
