package lib

import (
	"errors"
	"fmt"
)

// Reverse will reverse a given string
func Reverse(s string) (string, error) {
	strLen := len(s)

	// contrived error so we can test more things
	if strLen > 100 {
		errString := fmt.Sprintf(
			"while the library is really incredible it can only reverse string up to 100 character in length while the string provoided was %d characters in length",
			strLen,
		)
		return "", errors.New(errString)
	}

	// First off we should all agree that this is a very ugly way to reverse a string, right?
	// It is however very efficient as it only has to loop half the length
	// For `i` we count up from 0 and for `j` we count down from 4 (Ex: "hello")
	// So our first time around `i == "h"` and `j == "o"`
	// We flip these two and increment `i` (now 1) and decrement `j` (now 3): Result: "oellh"
	// We repeat the process `i == "e"` and `j == "l"`
	// We flip these two and increment `i` (2) and decrement `j` (2): Result: "olleh"
	// We do our check and you know what i (2) is NOT less than j (2) so we exit our loop
	// Having only iterated two times! Ugly... yes! Efficient... Also yes!
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// fmt.Printf("BEFORE: runes[i] == %c runes[j] == %c\n", runes[i], runes[j])
		runes[i], runes[j] = runes[j], runes[i]
		// fmt.Printf("AFTER: runes[i] == %c runes[j] == %c\n", runes[i], runes[j])
	}

	return string(runes), nil
}
