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

// Variable pour stocker l'état de jeu actuel

var currentGameState GameState

// Fonction pour gérer les requêtes de jeu

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

// Fonction pour gérer la création d'une nouvelle partie

func HandleNewGame(w http.ResponseWriter, r *http.Request) {
	difficulty := r.URL.Query().Get("difficulty") // Récupère le paramètre de difficulté

	wordsFilename := "words.txt"
	wordList, err := words.ReadWordsFromFile(wordsFilename)
	if err != nil {
		http.Error(w, "Error reading words from file", http.StatusInternalServerError)
		return
	}

	// Sélectionne un mot aléatoire à partir de la liste de mots
	wordToGuess := words.SelectRandomWord(wordList)
	guessedLetters := words.RevealLetters(wordToGuess, difficulty) // Initialiser avec des lettres révélées

	// Créer un nouvel état de jeu
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
	// Retourner l'état de jeu actuel
	json.NewEncoder(w).Encode(currentGameState)
}

// Fonction pour gérer les devinettes de l'utilisateur

func HandleUserGuess(w http.ResponseWriter, r *http.Request) {
	var userGuess struct {
		Guess string `json:"guess"`
	}
	// Décoder la devinette de l'utilisateur à partir de la requête JSON
	if err := json.NewDecoder(r.Body).Decode(&userGuess); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Récupérer la devinette de l'utilisateur
	guess := userGuess.Guess

	// Vérifier si l'entrée est une seule lettre minuscule
	if len(guess) != 1 || !strings.Contains("abcdefghijklmnopqrstuvwxyz", guess) {
		json.NewEncoder(w).Encode(currentGameState)
		return
	}
	// Convertir la devinette de l'utilisateur en rune
	guessedLetter := rune(guess[0])

	// Mettre à jour l'état de jeu actuel avec la lettre devinée
	if !strings.ContainsRune(currentGameState.GuessedLetters, guessedLetter) {
		currentGameState.GuessedLetters += string(guessedLetter)

		if !strings.ContainsRune(currentGameState.WordToGuess, guessedLetter) {
			currentGameState.Attempts--
		}
	}
	// Vérifier si le jeu est terminé
	if hasWon(currentGameState) {
		currentGameState.Status = "victoire"
	} else if currentGameState.Attempts <= 0 {
		currentGameState.Status = "défaite"
	}
	// Retourner l'état de jeu actuel
	json.NewEncoder(w).Encode(currentGameState)
}

// Fonction pour vérifier si le joueur a gagné

func hasWon(state GameState) bool {
	// Vérifie si toutes les lettres du mot à deviner ont été devinées
	for _, char := range state.WordToGuess {
		if char != ' ' && !strings.ContainsRune(state.GuessedLetters, char) {
			return false
		}
	}
	return true
}
