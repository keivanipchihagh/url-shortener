package shortener

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"

	"github.com/itchyny/base58-go"
)

const saltSize = 16

// Generate a random salt of fixed size
func generateSalt() []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}

	return salt
}

// Encode the input bytes to base58
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(encoded)
}

// Add a salt to the input and return the SHA256 hash
func toSha256(input string, salt []byte) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(salt)
}

func GenerateShortLink(originalUrl string) string {
	salt := generateSalt()
	urlHashBytes := toSha256(originalUrl, salt)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]
}
