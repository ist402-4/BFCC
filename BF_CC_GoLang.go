package main

import (
	"fmt"
	"strings"
	"unicode"
)

func decrypt(n int, ciphertext string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	ALPHABET := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := ""
	for _, l := range ciphertext {
		// iterate over string
		if unicode.IsUpper(l) {
			index := strings.Index(ALPHABET, string(l))
			i := (index - n + 26) % 26 //need to include +26 to keep the mod equation within range(0-25)
			result += string(ALPHABET[i])

		} else if unicode.IsLower(l) {
			index := strings.Index(alphabet, string(l))
			i := (index - n + 26) % 26 //need to include +26 to keep the mod equation within range(0-25)
			result += string(alphabet[i])
		} else {
			result += string(l) //for spaces
		}
	}
	return result
}

func main() {

	ciphertext := "Ugew gnwj zwjw Oslkgf"

	for key := 0; key < 26; key++ {
		dec := decrypt(key, ciphertext)
		fmt.Printf("%2d. %s\n", key, dec)
	}

}
