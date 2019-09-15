package piano_kata_p1

// hacky solution....
import "fmt"

func BlackOrWhiteKey(keyPressCount int) string {
	colour := ""
	counter := 1
	passes := 0

	startPat := []string{"white", "black", "white"}
	repeatingPat := []string{"white", "black", "white", "black", "white", "white", "black", "white", "black", "white", "black", "white"}

	for done := false; !done; {
		counter = counter % 88
		if counter == 0 {
			passes += 1
			counter = 1
		}

		if counter <= 3 {
			for pos, key := range startPat {

				if counter+(passes*88) == keyPressCount {
					colour = key
				}
				if pos+1 == len(startPat) {
					counter += 1
					break
				}
				counter += 1
			}
		}

		for pos, key := range repeatingPat {

			if counter+(passes*88) == keyPressCount {
				colour = key
			}

			if pos+1 == len(repeatingPat) {
				counter += 1
				break
			}

			if (counter%88 == 0) || (counter%88 == 88) {
				break
			}

			counter += 1
		}

		if counter+(passes*88) > 10000 {
			return "white"
		}

		if colour != "" {
			fmt.Sprintf("Completed %d!", keyPressCount)
			done = true
		}
	}
	return colour
}
