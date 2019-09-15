package century_from_year

import (
	"fmt"
	"github.com/DATA-DOG/godog"
)

func aYear() error {
	century(2000)
	return nil
}

func itShouldReturnOfThe(expected, year int) error {
	testResult := century(year)

	if testResult != expected {
		return fmt.Errorf("want %d; got %d", expected, testResult)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a year$`, aYear)
	s.Step(`^it should return (\d+) of the (\d+)$`, itShouldReturnOfThe)
}
