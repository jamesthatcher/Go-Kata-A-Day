package longest_vowel_chain

import "fmt"

// The vowel substrings in the word codewarriors are o,e,a,io.
// The longest of these has a length of 2. Given a lowercase string that
// has alphabetic characters only and no spaces, return the length of the
// longest vowel substring.

func Solve(s string) int {
	for pos, char := range s {
		fmt.Println(char, pos)
		fmt.Println(pos)
	}

	return 0
}
