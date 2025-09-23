package main 
import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"flag"
)

func encrypt(key string, message string, alphabet []rune, alpha map[rune]int) {

	cipherTxt := ""
	splitMsg := []rune(message) //split to iterate through individual letters
	splitKey := []rune(key) //see above
	
	k := 0

	for i:=0;i<len(message);i++{
		if (!strings.ContainsRune(" .-;:'+=*~`!@#$%^&*(){}?><,/|\\\"_", splitMsg[i])) {

			if k >= len(key){ //brings key back to first letter
				k = 0
			}
			
			P := alpha[splitMsg[i]] // P_i from The Formula
	
			K := alpha[splitKey[k]] // K_i from The Formula
	
			cipherTxt += string(alphabet[(P+K)%26])
			
			k += 1
		} else {
			cipherTxt += string(splitMsg[i])
		}
	}
	fmt.Println(cipherTxt)
}

func decrypt(key string, message string, alphabet []rune, alpha map[rune]int) {
	plainTxt := ""
	splitMsg := []rune(message) //split to iterate through individual letters
	splitKey := []rune(key) //see above
	
	k := 0

	for i:=0;i<len(message);i++{
		if (!strings.ContainsRune(" .-;:'+=*~`!@#$%^&*(){}?><,/|\\\"_", splitMsg[i])) {

			if k >= len(key){ //brings key back to first letter
				k = 0
			}
			
			C := alpha[splitMsg[i]] // C_i from The Formula
	
			K := alpha[splitKey[k]] // K_i from The Formula
	
			plainTxt += string(alphabet[(26+C-K)%26])
			
			k += 1
		} else {
			plainTxt += string(splitMsg[i])
		}
	}
	fmt.Println(plainTxt)
}

func vigenere(e bool, key string, message string) {
	//fmt.Println(d, e, key, message)

	alphabet := []rune("abcdefghijklmnopqrstuvwxyz") //rune array of letters to jump to index of a given letter
	alpha := make(map[rune]int) // map of the alphabet to get the index of a letter without searching the array

	for i := 0; i < 26; i++ { //fills the map
        letter := rune('a' + i)
        alpha[letter] = i
    }

	if e {
		encrypt(key, message, alphabet, alpha)
	} else {
		decrypt(key, message, alphabet, alpha)
	} 
}

func main() {

	//setting flags
	dCheck := flag.Bool("d", false, "Used to toggle the decode function")
	eCheck := flag.Bool("e", false, "Used to toggle the encode function")

	flag.Parse() //sorts flags from user input

	key := strings.TrimSpace(flag.Arg(0)) //gets first argument after the flags
	
	for *dCheck || *eCheck { //while either flag is true
		
		//receives message to encode
		input := bufio.NewReader(os.Stdin) //reads secondary user input
		msg, _ := input.ReadString('\n')
		msg = strings.TrimSpace(msg)
		
		vigenere(*eCheck, key, msg)

		//frees up input for another message
		input.Reset(os.Stdin)
	}
}
