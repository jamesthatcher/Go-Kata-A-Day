package abbreviate_a_two_word_name

import (
	"strings"
)

func AbbrevName(name string) string {
	split := strings.Split(name, " ")
	firstname := string(split[0][0])
	surname := string(split[1][0])
	return strings.ToUpper(firstname + "." + surname)
}
