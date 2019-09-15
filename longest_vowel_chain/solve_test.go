package longest_vowel_chain

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

func iNeedToCountTheConsecutiveVowels() error {
	return nil
}

func iCount(word string) error {
	Solve(word)
	return nil
}

func itShouldReturnOfLongestVowelChainIn(expected int, word string) error {
	testResult := Solve(word)
	if testResult != expected {
		return fmt.Errorf("want %d; got %d", expected, testResult)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I need to count the consecutive vowels$`, iNeedToCountTheConsecutiveVowels)
	s.Step(`^I count "([^"]*)"$`, iCount)
	s.Step(`^it should return (\d+) of longest vowel chain in "([^"]*)"$`, itShouldReturnOfLongestVowelChainIn)
}
