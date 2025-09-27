
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
)

const (
	address = "192.168.1.243" // since we are logging in on where its hosted
	// otherwise we would put the IP address of where we are logging in
	port = "21" // port number
	username = "anonymous"
	password = ""
	path = "/" // find by pwd
		   // where are the files you are interested in
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
	defer client.Close() // defer means this part executes last
	// so it executes at the end of main(). Its evaluated where it
	// is but executed at the end

	// read the listing of the files in the ftp server
	entries, err := client.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	// print out the file listing
	for _, entry := range entries {
		fmt.Printf("%s\t%s\n", entry.Mode().String(), entry.Name())
	}
}
