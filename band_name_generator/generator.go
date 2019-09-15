package band_name_generator

import "strings"

func bandNameGenerator(word string) string {
	runes := []rune(word)

	if string(runes[0]) == string(runes[len(word)-1]) {
		return strings.Title(string(runes) + string(runes[1:]))
	}
	return strings.Title("The " + word)
}
