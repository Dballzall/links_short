package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// sha256Of computes the SHA256 hash of the given input string.
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// base58Encoded converts the given bytes into a Base58 encoded string
// using the Bitcoin alphabet.
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		// If something goes wrong, print the error and exit.
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

// GenerateShortLink generates a short URL from the original URL and a userId.
// It works by concatenating the original link and userId, hashing them with SHA256,
// converting the hash into a big integer, encoding that as Base58, and taking the first 8 characters.
func GenerateShortLink(initialLink string, userId string) string {
	// Hash the concatenation of the original URL and userId.
	urlHashBytes := sha256Of(initialLink + userId)

	// Create a big integer from the hash bytes and then extract its Uint64 value.
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()

	// Format the number as a string, convert it to bytes, then Base58 encode it.
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))

	// Return the first 8 characters as the final short URL.
	return finalString[:8]
}
