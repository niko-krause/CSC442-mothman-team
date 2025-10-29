// main.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// keyPath is the filename to read the key from in the current directory.
const keyPath = "key"

// buffer size for streaming stdin/out (4KB)
const bufSize = 4096

func main() {
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
	keyIndex := 0 // position within key; if key shorter than input we'll repeat it

	for {
		n, err := in.Read(buf)
		if n > 0 {
			// XOR the bytes we read with the key (recycling key if needed)
			for i := 0; i < n; i++ {
				buf[i] ^= key[(keyIndex+i)%keyLen]
			}

			// advance keyIndex modulo keyLen
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
