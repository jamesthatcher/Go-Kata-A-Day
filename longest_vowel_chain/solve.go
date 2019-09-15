package longest_vowel_chain

import "regexp"

func Solve(s string) int {
	var n int
	r := regexp.MustCompile(`[aeiou]+`).FindAllString(s, -1)
	for _, v := range r {
		if len(v) > n {
			n = len(v)
		}
	}
	return n
}
