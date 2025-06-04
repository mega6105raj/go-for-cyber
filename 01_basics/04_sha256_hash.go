/* Documentation Section:
	This program generates the SHA-256 hash of a given input string. SHA-256 is a
cryptographic has function that produces a fixed-size(256-bit) output. It's widely used in:
	- Digital signatures
	- Integrity verification (file checksums)
	- Blockchain (Bitcoin uses SHA-256)
	- Password hashing (with salts)
This is one of the first steps toward cryptographically secure systems */

package main

import (
	"crypto/sha256" 	//SHA-256 hashing
	"encoding/hex"		//To display the hash in hexadecimal
	"fmt"	//For input/output
)

func main() {
	var input string
	fmt.Print("Enter a string to hash using SHA-256: ")
	fmt.Scanln(&input)

	hash := sha256.Sum256([]byte(input)) 	//Compute SHA-256 hash
	hashHex := hex.EncodeToString(hash[:])	//Convert byte array to hex string

	fmt.Println("SHA-256 Hash (hex): ", hashHex)	//Print the hex-encoded hash
}
