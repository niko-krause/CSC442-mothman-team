package main

// stegged-bit: -b -o1024
//stegged byte: -B -o1024 -i8 -> -B -o1025 -i2

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func store(bitFlag bool, byteFlag bool, offset int, interval int, wrapper string, hidden string){

	//the goal is to make this check if we're in bit or byte mode with bitFlag & byteFlag, then operate accordingly
	

	wStream, _ := os.Open(wrapper) //opens the wrapper file
	wFile, _ := wStream.Stat()
	wBytes := make([]byte, wFile.Size())
	bufio.NewReader(wStream).Read(wBytes) //reads file contents as bytes and writes them into wBytes

	sentinel := []byte {0, 255, 0, 0, 255, 0} //sentinel in decimal values

	hStream, _ := os.Open(hidden) //opens the file that we're hiding
	hFile, _ := hStream.Stat()
	hBytes := make([]byte, hFile.Size())
	bufio.NewReader(hStream).Read(hBytes) //reads file contents as bytes and writes them into hBytes

	j := 0 //counter for hidden file
	for i := offset; i<len(wBytes) && j<len(hBytes); i+=interval { //while we're in the limits of the wrapper and hidden file, iterating by the interval
		wBytes[i] = hBytes[j] //sets the wrapper's byte at index i to the next hidden file's byte
		j++

		if j >= len(hBytes){
			k := 0 //counter for the sentinel
			for i := i; i<len(wBytes) && k < len(sentinel); i+=interval {
				wBytes[i] = sentinel[k] //writes the sentinel to the wrapper file
				k++
			}

		}

	}
 
	os.WriteFile(wrapper, wBytes, 0777) //writes to file and gives it full permissions (find a way to do this without altering permissions)

	wStream.Close()
	hStream.Close()
}

func retrieve(bitFlag bool, byteFlag bool, offset int, interval int, wrapper string){

}

func main(){
	
	//setting flags
	sFlag := flag.Bool("s", false, "Used to toggle the store function")
	rFlag := flag.Bool("r", false, "Used to toggle the retrieve function")
	bitFlag := flag.Bool("b", false, "Toggles bit mode")
	byteFlag := flag.Bool("B", false, "Toggles byte mode")
	
	offset := flag.Int("o", 0,"offset value")
	interval := flag.Int("i", 1, "interval value") //optional


	wrapper := flag.String("w", "wrapper.txt", "stores the name of the wrapper file")
	hidden := flag.String("h", "golshi" , "stores the name of the hidden file") //optional
	
	flag.Parse()

	fmt.Println(*sFlag, *rFlag, *bitFlag, *byteFlag, *offset, *interval, *hidden, *wrapper) //prints all of the flag values for debugging

	if *sFlag {
		store(*bitFlag, *byteFlag, *offset, *interval, *wrapper, *hidden)
	} else if *rFlag {
		retrieve(*bitFlag, *byteFlag, *offset, *interval, *wrapper)
	}
}