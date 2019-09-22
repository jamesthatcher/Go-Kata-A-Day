package if_you_can_read_this

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

func aWord(str string) error {
	ToNato(str)
	return nil
}

func itShouldReturnACorrectTranslationOfThe(expected, str string) error {
	testResult := ToNato(str)

	if testResult != expected {
		return fmt.Errorf("want %s; got %s", expected, testResult)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a word "([^"]*)"$`, aWord)
	s.Step(`^it should return a correct "([^"]*)" translation of the "([^"]*)"$`, itShouldReturnACorrectTranslationOfThe)
}
