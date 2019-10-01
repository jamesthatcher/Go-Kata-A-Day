package rainfall

import (
	"math"
	"strconv"
	"strings"
)

func Mean(city string, data string) float64 {
	result := strings.Split(data, "\n")

	target := ""
	for _, c := range result {
		if strings.Split(c, ":")[0] == city {
			target = strings.Split(c, ":")[1]
		}
	}

	temps := strings.Split(target, ",")
	total := 0.0
	for _, t := range temps {
		temp, _ := strconv.ParseFloat(strings.Split(t, " ")[1], 64)
		total += float64(temp)
	}
	return total / float64(len(temps))
}
func Variance(city string, data string) float64 {
	avg := Mean(city, data)
	result := strings.Split(data, "\n")

	target := ""
	for _, c := range result {
		if strings.Split(c, ":")[0] == city {
			target = strings.Split(c, ":")[1]
		}
	}

	temps := strings.Split(target, ",")
	total := 0.0
	for _, t := range temps {
		temp, _ := strconv.ParseFloat(strings.Split(t, " ")[1], 64)
		diff := math.Pow(temp-avg, 2)
		total += float64(diff)
	}
	return total / float64(len(temps)-1)
}
