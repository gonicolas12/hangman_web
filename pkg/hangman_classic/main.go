package main

import (
	"encoding/json"
	"fmt"
	"hangman-classic/words"
	"log"
	"net/http"
	"strings"
)

type GameState struct {
	WordToGuess      string
	GuessedLetters   string
	Attempts         int
	HangmanPositions []string
}

var currentGameState GameState

func gameHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleNewGame(w, r)
	case "POST":
		handleUserGuess(w, r)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func handleNewGame(w http.ResponseWriter, r *http.Request) {
	wordsFilename := "words.txt"
	wordList, err := words.ReadWordsFromFile(wordsFilename)
	if err != nil {
		http.Error(w, "Error reading words from file", http.StatusInternalServerError)
		return
	}

	currentGameState = GameState{
		WordToGuess:    words.SelectRandomWord(wordList),
		GuessedLetters: "",
		Attempts:       10,
		HangmanPositions: []string{
			"hangman_steps/pendu 0.png",
			"hangman_steps/pendu 1.png",
			"hangman_steps/pendu 2.png",
			"hangman_steps/pendu 3.png",
			"hangman_steps/pendu 4.png",
			"hangman_steps/pendu 5.png",
			"hangman_steps/pendu 6.png",
			"hangman_steps/pendu 7.png",
			"hangman_steps/pendu 8.png",
			"hangman_steps/pendu 9.png",
			"hangman_steps/pendu 10.png",
		},
	}

	json.NewEncoder(w).Encode(currentGameState)
}

func handleUserGuess(w http.ResponseWriter, r *http.Request) {
	var userGuess struct {
		Guess string `json:"guess"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userGuess); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	guess := userGuess.Guess
	guessedLetter := rune(guess[0])

	if !strings.ContainsRune(currentGameState.GuessedLetters, guessedLetter) {
		currentGameState.GuessedLetters += string(guessedLetter)

		if !strings.ContainsRune(currentGameState.WordToGuess, guessedLetter) {
			currentGameState.Attempts--
		}
	}

	json.NewEncoder(w).Encode(currentGameState)
}

func main() {
	http.HandleFunc("/guess", gameHandler)
	http.Handle("/", http.FileServer(http.Dir("../../web/")))
	fmt.Println("Le serveur fonctionne sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
