package affine

import (
	"fmt"
	"unicode"
)

// https://en.wikipedia.org/wiki/Affine_cipher
// Coefficients found by brute force
const a int = 1
const b int = 21

// https://www.geeksforgeeks.org/implementation-affine-cipher/
func Decrypt(cipher string) (string, error) {
	aInverse := 0
	flag := 0
	msg := ""
	for i := 0; i < 26; i++ {
		flag = (a * i) % 26
		if flag == 1 {
			aInverse = i
		}
	}
	for i := 0; i < len(cipher); i++ {
		char := cipher[i]
		if char > unicode.MaxASCII {
			return "", fmt.Errorf("invalid char %s", string(char))
		}
		if unicode.IsSpace(rune(char)) {
			msg = msg + ""
		} else {
			plainChar := rune(((aInverse * (int(char) + 'A' - b) % 26) + 'A'))
			msg = msg + string(plainChar)
		}
	}
	return msg, nil
}
