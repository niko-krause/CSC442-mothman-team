package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func encrypt(key string, message string, alphabet []rune, alpha map[rune]int) string {

	cipherTxt := ""
	splitMsg := []rune(message) //split to iterate through individual letters
	splitKey := []rune(key)     //see above

	k := 0

	for i := 0; i < len(message); i++ {

		caps := false

		if unicode.IsLetter(splitMsg[i]) {//if the rune is a letter

			if k >= len(key) { //brings key back to first letter
				k = 0
			}

			if string(splitKey[k]) == " " {
				for string(splitKey[k]) == " "{
					k += 1
				}
			}

			if unicode.IsUpper(splitMsg[i]) { //checks to see if the rune is uppercase. If so, stores that for later and makes it lowercase
				caps = true
				splitMsg[i] = unicode.ToLower(splitMsg[i])
			}

			P := alpha[splitMsg[i]] // P_i from The Formula

			K := alpha[splitKey[k]] // K_i from The Formula

			if caps { //resolves stored capitalization
				cipherTxt += string(unicode.ToUpper(alphabet[(P+K)%26]))
			} else {
				cipherTxt += string(alphabet[(P+K)%26])
			}

			k += 1
		} else {
			cipherTxt += string(splitMsg[i])
		}
	}
	return cipherTxt
}

func decrypt(key string, message string, alphabet []rune, alpha map[rune]int) string {
	plainTxt := ""
	splitMsg := []rune(message) //split to iterate through individual letters
	splitKey := []rune(key)     //see above

	k := 0

	for i := 0; i < len(message); i++ {

		caps := false

		if unicode.IsLetter(splitMsg[i]) { //if the rune is a letter

			if k >= len(key) { //brings key back to first letter
				k = 0
			}

			if string(splitKey[k]) == " " {
				for string(splitKey[k]) == " "{
					k += 1
				}
			}

			if unicode.IsUpper(splitMsg[i]) { //checks to see if the rune is uppercase. If so, stores that for later and makes it lowercase
				caps = true
				splitMsg[i] = unicode.ToLower(splitMsg[i])
			}

			C := alpha[splitMsg[i]] // C_i from The Formula

			K := alpha[splitKey[k]] // K_i from The Formula

			if caps { //resolves stored capitalization
				plainTxt += string(unicode.ToUpper(alphabet[(26+C-K)%26]))
			} else {
				plainTxt += string(alphabet[(26+C-K)%26])
			}

			k += 1
		} else {
			plainTxt += string(splitMsg[i])
		}
	}
	return plainTxt
}

func vigenere(e bool, key string, message string) string {
	//fmt.Println(e, key, message)
	var output string

	alphabet := []rune("abcdefghijklmnopqrstuvwxyz") //rune array of letters to jump to index of a given letter
	alpha := make(map[rune]int)                      // map of the alphabet to get the index of a letter without searching the array

	for i := 0; i < 26; i++ { //fills the map
		letter := rune('a' + i)
		alpha[letter] = i
	}

	if e {
		output = encrypt(key, message, alphabet, alpha)
	} else {
		output = decrypt(key, message, alphabet, alpha)
	}
	return output
}

func main() {

	//setting flags
	dCheck := flag.Bool("d", false, "Used to toggle the decode function")
	eCheck := flag.Bool("e", false, "Used to toggle the encode function")

	flag.Parse() //sorts flags from user input

	key := strings.TrimSpace(flag.Arg(0)) //gets first argument after the flags
	key = strings.ToLower(key)            //forces the key to be lowercase (for ease of processing)
	if len(key) == 0 { //processes empty key like the letter a, leaving the messgae unencrypted
		key = "a"
	}
	
	for *dCheck || *eCheck { //while either flag is true

		//receives message to encode
		input := bufio.NewReader(os.Stdin) //reads secondary user input
		msg, _ := input.ReadString('\n')
		msg = strings.TrimSpace(msg)
		
		if len(msg) > 0 {
			newMsg := vigenere(*eCheck, key, msg)
			fmt.Println(newMsg)
		}

		//frees up input for another message
		input.Reset(os.Stdin)
	}
}
