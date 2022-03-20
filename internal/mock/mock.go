package mock

import (
	"math/rand"
	"strings"
)

func MockText(input string) string {
	// The function mock takes a string as an input and randomly changes the case of the individual letters.
	// The function returns the modified string
	var output string
	for _, char := range input {
		if rand.Intn(2) == 0 {
			output += strings.ToUpper(string(char))
		} else {
			output += strings.ToLower(string(char))
		}
	}
	return output
}
