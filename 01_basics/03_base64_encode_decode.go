/* Documentation Section:
	This program demonstrates how to Base64 encode and decode strings using Go.
Base64 encoding is sommonly used in cryptographic systems to:
	- Represent binary data in text (e.g., PEM keys, JWT tokens)
	- Transmit encoded ciphertext over protocols that expect ASCII
	- Embed binary blobs in JSON or XML
Understanding this is essential when working with keys, certificates, encrypted messages, or APIs */

package main

import (
	"encoding/base64"	//For Base64 encoding/decoding
	"fmt"	//For I/O
)

func main() {
	var input string
	fmt.Print("Enter a string to encode in Base64: ")
	fmt.Scanln(&input)

	// Convert string to Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(input))
	fmt.Println("Encoded (base64):", encoded)

	//Decode Base64 back to original string
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding Base64:", err)
		return
	}

	fmt.Println("Decoded string:", string(decodedBytes))
}
