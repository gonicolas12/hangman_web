package game

import (
	"encoding/json"
	"hangman-classic/words"
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

func GameHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		HandleNewGame(w, r)
	case "POST":
		HandleUserGuess(w, r)
	default:
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
	}
}

func HandleNewGame(w http.ResponseWriter, r *http.Request) {
	difficulty := r.URL.Query().Get("difficulty") // Nouveau: récupérer le paramètre de difficulté

	wordsFilename := "words.txt"
	wordList, err := words.ReadWordsFromFile(wordsFilename)
	if err != nil {
		http.Error(w, "Error reading words from file", http.StatusInternalServerError)
		return
	}

	wordToGuess := words.SelectRandomWord(wordList)
	guessedLetters := words.RevealLetters(wordToGuess, difficulty) // Modifié pour initialiser avec des lettres révélées

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

func HandleUserGuess(w http.ResponseWriter, r *http.Request) {
	var userGuess struct {
		Guess string `json:"guess"`
	}

	if err := json.NewDecoder(r.Body).Decode(&userGuess); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	guess := userGuess.Guess

	// Vérifier si l'entrée est une seule lettre minuscule
	if len(guess) != 1 || !strings.Contains("abcdefghijklmnopqrstuvwxyz", guess) {
		// Si non, simplement retourner l'état actuel sans modifier les tentatives ou les lettres devinées
		json.NewEncoder(w).Encode(currentGameState)
		return
	}

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
