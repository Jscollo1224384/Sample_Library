package wordutils

import (
	"testing"
	"unicode/utf8"
)

func TestWordJumble(t *testing.T) {
	word := "example"
	jumbled := WordJumble(word)
	if len(jumbled) != len(word) {
		t.Errorf("Jumbled word length mismatch: got %d, want %d", len(jumbled), len(word))
	}
	// Check all original runes are present
	for _, r := range word {
		if countRune(jumbled, r) != countRune(word, r) {
			t.Errorf("Rune %q count mismatch in jumbled word", r)
		}
	}
}

func countRune(s string, r rune) int {
	count := 0
	for _, x := range s {
		if x == r {
			count++
		}
	}
	return count
}

func TestWordCypher(t *testing.T) {
	cases := []struct {
		in     string
		shift  int
		expect string
	}{
		{"abc", 1, "bcd"},
		{"xyz", 2, "zab"},
		{"ABC", 3, "DEF"},
		{"Hello, World!", 5, "Mjqqt, Btwqi!"},
		{"Go123", 1, "Hp123"},
		{"Zebra", -1, "Ydaqz"},
	}
	for _, c := range cases {
		got := WordCypher(c.in, c.shift)
		if got != c.expect {
			t.Errorf("WordCypher(%q, %d) = %q; want %q", c.in, c.shift, got, c.expect)
		}
	}
}

func TestWordJumbleUnicode(t *testing.T) {
	word := "héllo"
	jumbled := WordJumble(word)
	if utf8.RuneCountInString(jumbled) != utf8.RuneCountInString(word) {
		t.Errorf("Jumbled unicode word length mismatch: got %d, want %d", utf8.RuneCountInString(jumbled), utf8.RuneCountInString(word))
	}
}
