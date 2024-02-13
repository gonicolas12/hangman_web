package words

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

// Reads the words from the given file and returns them as a slice of strings

func ReadWordsFromFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(content), "\n")
	return words, nil
}

// Select a random word from the given slice of strings

func SelectRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

// Reveal a number of letters from the given word based on the given difficulty

func RevealLetters(word string, difficulty string) string {
	var n int
	switch difficulty {
	case "1": // Facile
		n = len(word)/2 - 1
	case "2": // Moyen
		n = len(word)/2 - 2
	case "3": // Difficile
		n = 1
	case "4": // Extrême
		return "" // Ne révélez aucune lettre
	default:
		return ""
	}

	// Logique pour révéler `n` lettres uniques
	revealedLetters := ""
	for i := 0; i < n && i < len(word); i++ {
		letter := string(word[rand.Intn(len(word))])
		if !strings.Contains(revealedLetters, letter) {
			revealedLetters += letter
		} else {
			i-- // Réessayer si la lettre est déjà choisie
		}
	}
	return revealedLetters
}
