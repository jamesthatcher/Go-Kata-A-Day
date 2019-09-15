package band_name_generator

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

func iHaveWithoutTheStartAndEndMatching(name string) error {
	bandNameGenerator(name)
	return nil
}

func iGenerate() error {
	bandNameGenerator("Cheese")
	return nil
}

func itShouldReturn(name, expected string) error {
	testResult := bandNameGenerator(name)

	if testResult != expected {
		return fmt.Errorf("want %s; got %s", expected, testResult)
	}
	return nil
}

func iHaveWithTheStartAndEndMatching(name string) error {
	bandNameGenerator("Cheese")
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have "([^"]*)" without the start and end matching$`, iHaveWithoutTheStartAndEndMatching)
	s.Step(`^I generate$`, iGenerate)
	s.Step(`^it should return "([^"]*)" of the "([^"]*)"$`, itShouldReturn)
	s.Step(`^I have "([^"]*)" with the start and end matching$`, iHaveWithTheStartAndEndMatching)
}
