package words

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
)

// Print the welcome message in ascii art

func PrintAsciiArtWelcome() {
	// hangman in ascii art
	fmt.Println("  _   _                                          \n | | | | __ _ _ __   __ _ _ __ ___   __ _ _ __  \n | |_| |/ _` | '_ \\ / _` | '_ ` _ \\ / _` | '_ \\ ")
	fmt.Println(" |  _  | (_| | | | | (_| | | | | | | (_| | | | |")
	fmt.Println(" |_| |_|\\__,_|_| |_|\\__, |_| |_| |_|\\__,_|_| |_|")
	fmt.Println("                    |___/                       ")
	// game in ascii art
	fmt.Println("  ____    ")
	fmt.Println(" / ___| __ _ _ __ ___   ___  ")
	fmt.Println("| |  _ / _` | '_ ` _ \\ / _ \\")
	fmt.Println("| |_| | (_| | | | | | |  __/")
	fmt.Println(" \\____|\\__,_|_| |_| |_|\\___|")
}

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

// Display he word to guess that have been guessed so far

func DisplayWord(word, guessedLetters string) string {
	displayedWord := ""
	for _, char := range word {
		if strings.ContainsRune(guessedLetters, char) {
			displayedWord += string(char) + " "
		} else {
			displayedWord += "_ "
		}
	}
	return displayedWord
}
