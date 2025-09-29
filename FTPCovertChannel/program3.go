
/*
 author: Chandler
 date: 24th Sept, 2025
 description: go code to log into an ftp server, and get the file listings
*/

package main

import(
	"fmt"
	"log"
	"github.com/secsy/goftp"
	//"program3.go/BinaryDecoder"
	"strconv" 
)

const (
	address = "138.47.144.148" // since we are logging in on where its hosted
	// otherwise we would put the IP address of where we are logging in
	port = "21" // port number
	username = "anonymous"
	password = ""
	path = "/" // find by pwd
		   // where are the files you are interested in
	// variable to change the amount of bits that we are reading (7 or 10)
	METHOD = 7
)

func main() {
	// update the config struct that is part of the goftp library
	config := goftp.Config{
		User: username,
		Password: password,
		//ActiveTransfers: true,
	}

	// login to the ftp server
	client, err := goftp.DialConfig(config, address + ":" + port)

	if err != nil {// nil is the go version of null or none
		log.Fatal(err)
	}
	//defer client.Close() // defer means this part executes last
	// so it executes at the end of main(). Its evaluated where it
	// is but executed at the end

	// read the listing of the files in the ftp server
	entries, err := client.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// print out the file listing
	var binarystring string

	for _, entry := range entries {
		//fmt.Printf("%s\t%s\n", entry.Mode().String(), entry.Name())
		// take the result from entry.Mode() turn into string 
		//stringRep := entry.Mode()
		
		// turn that string into binary
		//binaryPerms := strconv.FormatUint(stringRep, 2)
		//fmt.Println( entry.Mode())
		bitstring:= fmt.Sprintf("%010b", entry.Mode().Perm()) // used throw away var here

		// if the bit amount is 7 need to account for noise and skip it if it shows up
		if METHOD == 7 {
			if bitstring[0] == '1' || bitstring[1] == '1' || bitstring[2] == '1' {
				continue // skippa 
			}

			// add only the last 7 bits into the string 
			binarystring += bitstring[3:]
		}else if METHOD == 10 {
			binarystring += bitstring 
		}
	}
	// temp debug print 
	//fmt.Print(binarystring)
	//fmt.Println()

	// disconnecting from FTP server 
	client.Close()

	// now that the FTP server has been disconnected from, we isolate and decode the file permissions
	result := Decode(binarystring)
	fmt.Print(result)


}

// stole my own code from the program1
func Decode(inputString string) string{ // needs to now return a string 
	// taking the input string as parameter, decode, return decodedString as result
	errorString := "Input cannot be interpreted"
	var decodedString string  
	//fmt.Print("Enter a binary string: ") 
	//fmt.Scan(&input) // taking user input 

	//fmt.Printf("Your input was: %s", input) // printing user input back at them
	//fmt.Println() 

	// establishing default group size (here its always 7 bit ASCII)
	grouping := 7
	
	// check if user input is a multiple of 7 and if not then pad as needed
	// if user input is a multiple of 7 proceed as normal
	if len(inputString) % grouping != 0 {
		padAmount := grouping - (len(inputString) %  grouping) // padding by the differnece between group amount and the input length
		for i := 0; i < padAmount; i++{
			// pad by adding zeros at the end 
			inputString = inputString + "0"
		}
	}

	// split user input into groups of 7 and then feed those in to be interpreted into ASCII

	// grouping the total input into the grouping size  
	for i := 0; i < len(inputString); i += grouping {
		inputGroup := inputString[i : i + grouping]
		
		// converting input into integer then that is feed into ascii to become a character 

		val, err := strconv.ParseInt(inputGroup, 2, 64) // interpreting user input as binary then translating its value to dec
		// if the input cannot be interpreted then throw an error 
		if err != nil{
			fmt.Println("Error: ", err) 
			return errorString
		}

		// concatenating those characters to form a final output string 
		// printing the groups by formatting as chars and not printing a newline
		//fmt.Printf("%c", val)
		decodedString += fmt.Sprintf("%c", val)

	}
	return decodedString
}

