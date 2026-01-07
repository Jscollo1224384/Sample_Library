// Package wordutils provides utilities for word manipulation.
package wordutils

import (
	"math/rand"
	"strings"
	"time"
)

// WordJumble returns a new string with the letters of the input word randomly shuffled.
func WordJumble(word string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	runes := []rune(word)
	n := len(runes)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// WordCypher returns a new string with each letter shifted by the given offset using a Caesar cipher.
// Only letters (A-Z, a-z) are shifted; other characters remain unchanged.
func WordCypher(word string, shift int) string {
	shift = shift % 26
	var sb strings.Builder
	for _, r := range word {
		switch {
		case r >= 'a' && r <= 'z':
			sb.WriteRune('a' + (r-'a'+rune(shift)+26)%26)
		case r >= 'A' && r <= 'Z':
			sb.WriteRune('A' + (r-'A'+rune(shift)+26)%26)
		default:
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

// ReverseWord returns a new string with the characters of the input word in reverse order.
func ReverseWord(word string) string {
	runes := []rune(word)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// RemoveSpaces returns a new string with all spaces removed from the input word.
// Example:
//
//	RemoveSpaces("hello world") // returns "helloworld"
//	RemoveSpaces(" Go  Lang ") // returns "GoLang"
func RemoveSpaces(word string) string {
	return strings.ReplaceAll(word, " ", "")
}
