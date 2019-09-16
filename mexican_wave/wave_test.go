package mexican_wave

import (
	"fmt"
	"github.com/DATA-DOG/godog"
	"strings"
)

func aLowerCase(input string) error {
	testResult := strings.ToLower(input)

	if testResult != input {
		return fmt.Errorf("want %s; got %s", input, testResult)
	}
	return nil
}

func transformIntoAWave() error {
	return nil
}

func itShouldReturnOfEach(expected, input string) error {
	testResult := strings.Join(wave(input), " ")
	if testResult != expected {
		return fmt.Errorf("want %s; got %s", expected, testResult)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a lower case "([^"]*)"$`, aLowerCase)
	s.Step(`^transform into a wave$`, transformIntoAWave)
	s.Step(`^it should return "([^"]*)" of each "([^"]*)"$`, itShouldReturnOfEach)
}
