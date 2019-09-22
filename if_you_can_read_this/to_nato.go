package if_you_can_read_this

import (
	"strings"
)

func ToNato(str string) string {
	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "X-ray", "Yankee", "Zulu"}

	var match []string
	for _, s := range strings.Split(str, "") {
		if strings.ContainsAny(s, "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~") {
			match = append(match, s)
		}

		for _, f := range nato {
			it := string(f[0])
			if strings.ToLower(s) == strings.ToLower(it) {
				match = append(match, f)
			}
		}

	}
	return strings.Join(match, " ")
}
