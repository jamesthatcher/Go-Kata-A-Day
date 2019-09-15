package piano_kata_p1

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

func iStoppedAt(key int) error {
	BlackOrWhiteKey(key)
	return nil
}

func itShouldReturnTheOfThe(expected string, key int) error {
	testResult := BlackOrWhiteKey(key)
	if testResult != expected {
		return fmt.Errorf("want %s; got %s", expected, testResult)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I stopped at (\d+)$`, iStoppedAt)
	s.Step(`^it should return the "([^"]*)" of the (\d+)$`, itShouldReturnTheOfThe)
}
