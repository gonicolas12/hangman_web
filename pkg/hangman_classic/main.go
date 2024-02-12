package main

import (
	"encoding/json"
	"fmt"
	"hangman-classic/words"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type GameState struct {
	Difficulty       string
	WordToGuess      string
	GuessedLetters   string
	Attempts         int
	HangmanPositions []string
	Status           string
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
	difficulty := r.URL.Query().Get("difficulty") // Nouveau: récupérer le paramètre de difficulté

	wordsFilename := "words.txt"
	wordList, err := words.ReadWordsFromFile(wordsFilename)
	if err != nil {
		http.Error(w, "Error reading words from file", http.StatusInternalServerError)
		return
	}

	wordToGuess := words.SelectRandomWord(wordList)
	guessedLetters := revealLetters(wordToGuess, difficulty) // Modifié pour initialiser avec des lettres révélées

	currentGameState = GameState{
		Difficulty:     difficulty,
		WordToGuess:    wordToGuess,
		GuessedLetters: guessedLetters,
		Attempts:       10,
		HangmanPositions: []string{
			"hangman_steps/pendu0.png",
			"hangman_steps/pendu1.png",
			"hangman_steps/pendu2.png",
			"hangman_steps/pendu3.png",
			"hangman_steps/pendu4.png",
			"hangman_steps/pendu5.png",
			"hangman_steps/pendu6.png",
			"hangman_steps/pendu7.png",
			"hangman_steps/pendu8.png",
			"hangman_steps/pendu9.png",
			"hangman_steps/pendu10.png",
		},
		Status: "en cours",
	}

	json.NewEncoder(w).Encode(currentGameState)
}

func revealLetters(word string, difficulty string) string {
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

	if hasWon(currentGameState) {
		currentGameState.Status = "victoire"
	} else if currentGameState.Attempts <= 0 {
		currentGameState.Status = "défaite"
	}

	json.NewEncoder(w).Encode(currentGameState)
}

func hasWon(state GameState) bool {
	for _, char := range state.WordToGuess {
		if char != ' ' && !strings.ContainsRune(state.GuessedLetters, char) {
			return false
		}
	}
	return true
}

func main() {
	http.HandleFunc("/guess", gameHandler)
	http.Handle("/", http.FileServer(http.Dir("../../web/templates/")))
	fmt.Println("Le serveur fonctionne sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
