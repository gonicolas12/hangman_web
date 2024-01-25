package game

import "math/rand"

// Check if the given map contains the given rune

func Contains(runes map[rune]bool, r rune) bool {
	_, found := runes[r]
	return found
}

// Reveal the letters of the word that have been guessed so far

func RevealLetters(word string, guessedLetters map[rune]bool) string {
	revealedWord := ""
	for _, char := range word {
		if Contains(guessedLetters, char) {
			revealedWord += string(char) + " "
		} else {
			revealedWord += "_ "
		}
	}
	return revealedWord
}

// Calculates the number of letters to reveal at the start of the game based on word length

func InitialRevealedLetters(wordLength int) int {
	switch {
	case wordLength <= 5:
		return 1
	case wordLength <= 7:
		return 2
	case wordLength <= 9:
		return 3
	case wordLength <= 11:
		return 4
	default:
		return 5
	}
}

// Reveal n random letters of the word

func RevealRandomLetters(word string, n int, guessedLetters map[rune]bool) {
	// Create a slice to hold unique characters from the word
	uniqueChars := make([]rune, 0)
	for _, char := range word {
		if !guessedLetters[char] {
			uniqueChars = append(uniqueChars, char)
		}
	}

	// Shuffle the slice of unique characters
	rand.Shuffle(len(uniqueChars), func(i, j int) {
		uniqueChars[i], uniqueChars[j] = uniqueChars[j], uniqueChars[i]
	})

	// Reveal up to 'n' unique characters from the shuffled slice
	for i := 0; i < n && i < len(uniqueChars); i++ {
		guessedLetters[uniqueChars[i]] = true
	}
}
