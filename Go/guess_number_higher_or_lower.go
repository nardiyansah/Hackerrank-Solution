package main

func guess(lower int) int {
	return 0
}

func guessNumber(n int) int {
	lower, upper := 0, n

	if guess(lower) == 0 {
		return lower
	}

	if guess(upper) == 0 {
		return upper
	}

	for {
		guessedNumber := (lower + upper) / 2
		resultGuess := guess(guessedNumber)
		if resultGuess == 0 {
			return guessedNumber
		} else if resultGuess == -1 {
			upper = guessedNumber
		} else {
			lower = guessedNumber
		}

	}
}
