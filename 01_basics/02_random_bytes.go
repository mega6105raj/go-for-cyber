/*Documentation Section :
	This program generates a specified number of cryptographically secure random bytes using Go's crypto/rand package.
Unlike math/rand, crypto/rand is designed to be secure and unpredicatble, ideal for cryptographic contexts
*/
package main

import (
	"crypto/rand" //For secure random number generation
	"encoding/hex" //To print the random bytes in hexadecimal
	"fmt"	//For standdard input/output
	"log" 	//For error logging
)

func main () {
	const size = 16 	//Number of random bytes to generate - 128 bits

	randomBytes := make([]byte, size)  //Creates a byte slice of desired length

	_, err := rand.Read(randomBytes)	//Fill the slice with secure random bytes
	if err != nil {
		log.Fatal("Error generating random bytes:", err)	//Log and exit if error
	}

	fmt.Println("Random Bytes (hex):", hex.EncodeToString(randomBytes))	//Print in hex
}
