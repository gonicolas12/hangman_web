package words

import (
	"io/ioutil"
	"math/rand"
	"strings"
)

// Lit les mots à partir d'un fichier et les retourne sous forme de tranche de chaînes

func ReadWordsFromFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	words := strings.Split(string(content), "\n")
	return words, nil
}

// Sélectionne un mot aléatoire à partir de la liste de mots

func SelectRandomWord(words []string) string {
	return words[rand.Intn(len(words))]
}

// Révèle un certain nombre de lettres du mot en fonction de la difficulté

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

	// Logique pour révéler 'n' lettres uniques
	revealedLetters := ""
	for i := 0; i < n && i < len(word); i++ {
		letter := string(word[rand.Intn(len(word))])
		if !strings.Contains(revealedLetters, letter) {
			// Si la lettre n'est pas déjà révélée, ajoutez-la à la liste des lettres révélées
			revealedLetters += letter
		} else {
			i-- // Réessayer si la lettre est déjà choisie
		}
	}
	// Remplacez les lettres révélées par des tirets dans le mot
	return revealedLetters
}
