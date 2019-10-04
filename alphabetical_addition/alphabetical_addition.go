package alphabetical_addition

func AddLetters(letters []rune) rune {
	var sum = 0
	for _, x := range letters {
		sum += int(x) - 96
	}
	return rune("zabcdefghijklmnopqrstuvwxyz"[sum%26])
}
