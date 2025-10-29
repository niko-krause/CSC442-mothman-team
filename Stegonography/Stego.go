/*
NOTICE: flags are parsed by the program when written with an equals sign (i.e. -i=2424). Otherwise, the execution of this program is unchanged.
*/



package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"log"
)

func store(bitFlag bool, byteFlag bool, offset int, interval int, wrapper string, hidden string){

	//the goal is to make this check if we're in bit or byte mode with bitFlag & byteFlag, then operate accordingly
	
	wStream, err := os.Open(wrapper) //opens the wrapper file
	if err != nil {
		log.Fatal(err)
	}
	wFile, _ := wStream.Stat()
	wBytes := make([]byte, wFile.Size())
	bufio.NewReader(wStream).Read(wBytes) //reads file contents as bytes and writes them into wBytes
	

	sentinel := []byte {0, 255, 0, 0, 255, 0} //sentinel in decimal values

	hStream, err2 := os.Open(hidden) //opens the file that we're hiding
	if err2 != nil {
		log.Fatal(err)
	}
	hFile, _ := hStream.Stat()
	hBytes := make([]byte, hFile.Size())
	bufio.NewReader(hStream).Read(hBytes) //reads file contents as bytes and writes them into hBytes

	//appends the sentinel to the hidden file's byte array
	for _, thing := range sentinel {
		hBytes = append(hBytes, thing)
	}

	if (bitFlag){
		wIndex := offset
		for i := 0; i < len(hBytes); i++ { //for every byte in the hidden file's byte array
			for j := 0; j<8; j++{
				wBytes[wIndex] &= 0b11111110
				wBytes[wIndex] |= ((hBytes[i] & 0b10000000) >> 7)
				hBytes[i] <<= 1
				wIndex += interval
			}
		}	
	} else if byteFlag {
		j := 0 //counter for hidden file
		for i := offset; i<len(wBytes) && j<len(hBytes); i+=interval { //while we're in the limits of the wrapper and hidden file, iterating by the interval
	
			wBytes[i] = hBytes[j] //sets the wrapper's byte at index i to the next hidden file's byte
			j++
		}

		
	}
 
	os.WriteFile(wrapper, wBytes, 0777) //writes to file and gives it full permissions (find a way to do this without altering permissions)

	wStream.Close()
	hStream.Close()
}

func retrieve(bitFlag bool, byteFlag bool, offset int, interval int, wrapper string){
	
	wStream, err := os.Open(wrapper) //opens the wrapper file
	if err != nil {
		log.Fatal(err)
	}
	wFile, _ := wStream.Stat()
	wBytes := make([]byte, wFile.Size())
	bufio.NewReader(wStream).Read(wBytes) //reads file contents as bytes and writes them into wBytes

	sentinel := []byte {0, 255, 0, 0, 255, 0} //sentinel in decimal values

	hBytes := []byte {}
	endcoming := []byte {} //stores suspected sentinel bytes for comparison

	if (byteFlag) {
		for i := offset; i<len(wBytes); i+=interval { //while we're in the limits of the wrapper and hidden file, iterating by the interval
	
			hBytes = append(hBytes, wBytes[i]) //appends the suspected hidden byte to the proper array
			
			endcoming = append(endcoming, wBytes[i])
			
			if slices.Equal(endcoming, sentinel[0:len(endcoming)]) { //if endcoming and the sentinel are currently equal
				if (len(endcoming) == 6){ //if the sentinel is the same length as the sentinel checker
					hBytes = slices.Delete(hBytes, len(hBytes)-6, len(hBytes)) //removes the last 6 bytes, thus removing the sentinel
					fmt.Print(string(hBytes))
					break
				} 
			} else {
				endcoming = nil
			}
		}
	} else if (bitFlag) {

		newBytes := []byte {}

		wIndex := offset
		for wIndex < len(wBytes) {
			tempByte := byte(0)

			for j := 0; j<8 && wIndex < len(wBytes); j++{ 
				tempByte <<= 1 //moves the bits over to the left by 1
				tempByte |= ((wBytes[wIndex] & 0b00000001)) //does the or operation with the least significant bit of wBytes, setting them equal
				wIndex += interval
			}

			hBytes = append(hBytes, tempByte)
		}

		

		for i:=0; i<len(hBytes);i++ {
			
			hByte := byte(hBytes[i])

			newBytes = append(newBytes, hByte)
			
			
			endcoming = append(endcoming, hByte)
			
			if (slices.Equal(endcoming, sentinel[0:len(endcoming)])){ 
				if len(endcoming) == 6 {
					newBytes = slices.Delete(newBytes, len(newBytes)-6, len(newBytes)) //removes the last 6 bytes, thus removing the sentinel
					fmt.Print(string(newBytes))
					break
				}
			} else {
				endcoming = nil
			}
			
		
		}

	}

	
	wStream.Close()
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

	//fmt.Println(*sFlag, *rFlag, *bitFlag, *byteFlag, *offset, *interval, *hidden, *wrapper) //prints all of the flag values for debugging

	if *sFlag {
		store(*bitFlag, *byteFlag, *offset, *interval, *wrapper, *hidden)
	} else if *rFlag {
		retrieve(*bitFlag, *byteFlag, *offset, *interval, *wrapper)
	}
}