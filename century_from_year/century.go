package century_from_year

func century(year int) int {
	if year%100 != 0 {
		year += 100
	}
	return year / 100
}
