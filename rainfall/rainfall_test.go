package rainfall

import "github.com/DATA-DOG/godog"

type meanWeather struct {
}

type varWeather struct {
}

func iHaveDataOfTypeAndCity(data, city string) error {
	return nil
}

func theMeanIsCalculated() error {
	//Mean(c, w)
	return nil
}

func iShouldHaveAnAverageRainfallOf(arg1 int) error {
	return godog.ErrPending
}

func theVarianceIsCalculated() error {
	return godog.ErrPending
}

func iShouldHaveAnVarianceRainfallOf(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have data of type "([^"]*)" and city "([^"]*)"$`, iHaveDataOfTypeAndCity)
	s.Step(`^the mean is calculated$`, theMeanIsCalculated)
	s.Step(`^I should have an average rainfall of (\d+)\.(\d+)$`, iShouldHaveAnAverageRainfallOf)
	s.Step(`^I should have an average rainfall of -(\d+)$`, iShouldHaveAnAverageRainfallOf)
	s.Step(`^the variance is calculated$`, theVarianceIsCalculated)
	s.Step(`^I should have an variance rainfall of (\d+)\.(\d+)$`, iShouldHaveAnVarianceRainfallOf)
	s.Step(`^I should have an variance rainfall of -(\d+)$`, iShouldHaveAnVarianceRainfallOf)
}
