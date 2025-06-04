/* Documenttaion Section:
	This Go program computes the SHA-512 hash of a user-provided message.
SHA-512 is part of the SHA-2 family and produces a 512-bit (64-byte) digest.
	- Used in: Cryptographic signatures, password hashing, TLS, blockchain
	- Secure: Comapres to SHA-1, SHA-512 has stronger resistance against collisiona nd preimage attacks
*/

package main

import (
	"crypto/sha512" 	//For SHA-512 hash function
	"encoding/hex" 	//For converting bytes to hex
	"fmt"
)

func main() {
	var message string
	fmt.Print("Enter a message to hash with SHA-512: ")
	fmt.Scanln(&message)

	hash := sha512.New()	//Creates a new SHA-512 hasher
	hash.Write([]byte(message))	//Write the message bytes
	hashBytes := hash.Sum(nil)	//Compute the final hash
	hashHex := hex.EncodeToString(hashBytes) 	//Convert to hex string

	fmt.Println("SHA-512 Hash (hex):", hashHex)
}
