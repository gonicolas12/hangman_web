package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hangman-classic/game"
	"hangman-classic/hangman"
	"hangman-classic/words"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// GameState holds the state of the game

type GameState struct {
	WordToGuess      string
	GuessedLetters   map[rune]bool
	Attempts         int
	HangmanPositions []string
}

// ANSI color codes

const (
	ColorGreen  = "\033[32m"
	ColorBlue   = "\033[34m"
	ColorOrange = "\033[33m"
	ColorRed    = "\033[31m"
	ColorReset  = "\033[0m"
)

func main() {
	words.PrintAsciiArtWelcome()
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go words.txt hangman.txt --startWith save.txt")
		return
	}
	// Read the words from the file
	wordsFilename := os.Args[1]
	hangmanFilename := os.Args[2]
	wordList, err := words.ReadWordsFromFile(wordsFilename)
	if err != nil {
		fmt.Printf("Error reading words from file: %v\n", err)
		return
	}
	// Read the hangman positions from the file
	hangmanPositions, err := hangman.ReadHangmanPositionsFromFile(hangmanFilename)
	if err != nil {
		fmt.Printf("Error reading hangman positions from file: %v\n", err)
		return
	}
	// Reverse the hangman positions
	hangman.ReverseHangmanPositions(hangmanPositions)
	var gameState GameState
	if len(os.Args) > 3 && os.Args[3] == "--startWith" {
		savedGame, err := ioutil.ReadFile(os.Args[4])
		if err != nil {
			fmt.Println("\nError reading saved game:", err)
			return
		}
		err = json.Unmarshal(savedGame, &gameState)
		if err != nil {
			fmt.Println("Error decoding saved game:", err)
			return
		}
		fmt.Println("\nWelcome Back, you have", gameState.Attempts, "attempts remaining.")
		hangman.DisplayHangman(gameState.Attempts, gameState.HangmanPositions)
	} else {
		// Create a new game
		gameState = GameState{
			WordToGuess:      words.SelectRandomWord(wordList),
			GuessedLetters:   make(map[rune]bool),
			Attempts:         11, // 11 instead of 10 to avoid a bug
			HangmanPositions: hangmanPositions,
		}
		// Ask the user to select the difficulty level
		fmt.Print("\nSelect Difficulty Level:\n")
		fmt.Printf("1: %sEasy%s | ", ColorGreen, ColorReset)
		fmt.Printf("2: %sAverage%s | ", ColorBlue, ColorReset)
		fmt.Printf("3: %sHard%s | ", ColorOrange, ColorReset)
		fmt.Printf("4: %sExtreme%s\n", ColorRed, ColorReset)
		var difficultyLevel int
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Enter difficulty (1-4): ")
			difficultyInput, _ := reader.ReadString('\n')
			difficultyInput = strings.TrimSpace(difficultyInput)
			var err error
			difficultyLevel, err = strconv.Atoi(difficultyInput)
			if err != nil || difficultyLevel < 1 || difficultyLevel > 4 {
				fmt.Println("Invalid input. Please enter a number between 1 and 4.")
				continue
			}
			break
		}
		switch difficultyLevel {
		case 2:
			initialRevealedLetters := game.InitialRevealedLetters(len(gameState.WordToGuess)) - 1
			game.RevealRandomLetters(gameState.WordToGuess, initialRevealedLetters, gameState.GuessedLetters)
		case 3:
			game.RevealRandomLetters(gameState.WordToGuess, 1, gameState.GuessedLetters)
		case 4:
			// No letters revealed for extreme difficulty
		default:
			initialRevealedLetters := game.InitialRevealedLetters(len(gameState.WordToGuess))
			game.RevealRandomLetters(gameState.WordToGuess, initialRevealedLetters, gameState.GuessedLetters)
		}
		fmt.Println("\nGood Luck, you have 10 attempts.")
	}
	// Start the game
	for gameState.Attempts > 0 {
		fmt.Println("Word to guess:", game.RevealLetters(gameState.WordToGuess, gameState.GuessedLetters))
		fmt.Print("Choose: ")
		// Read the user input
		reader := bufio.NewReader(os.Stdin)
		guess, _ := reader.ReadString('\n')
		guess = strings.TrimSpace(guess)
		if guess == "" {
			fmt.Println("No input provided. Please enter a letter or a word.")
			continue
		}
		guessedWord := len(guess) > 1
		// New condition: Check if the guess contains only letters when it's more than one character long
		isOnlyLetters := guessedWord && regexp.MustCompile("^[a-zA-Z]+$").MatchString(guess)
		// Check if the input is a single letter and if it's a lowercase letter
		if !guessedWord && (len(guess) != 1 || !('a' <= guess[0] && guess[0] <= 'z')) {
			fmt.Println("Invalid input. Please enter a single lowercase letter.")
			continue
		}
		// Check for a valid word guess (more than one character and only letters)
		if guessedWord && !isOnlyLetters {
			fmt.Println("Invalid guess. Please enter only letters to guess the word.")
			continue
		}
		// Check if the user wants to stop the game
		if guess == "STOP" {
			saveData, err := json.Marshal(gameState)
			if err != nil {
				fmt.Println("Error saving game:", err)
				continue
			}
			err = ioutil.WriteFile("save.txt", saveData, 0644)
			if err != nil {
				fmt.Println("Error writing save file:", err)
				continue
			}
			fmt.Println("Game Saved in save.txt.")
			return
		}
		// Check if the user guessed the word
		if guessedWord {
			if guess == gameState.WordToGuess {
				fmt.Println("Congratulations! You've guessed the word correctly!")
				return
			} else {
				fmt.Println("Incorrect guess. You lose 2 attempts.")
				gameState.Attempts -= 2
				if gameState.Attempts <= 0 {
					fmt.Println("Haha, you lost! The word was:", gameState.WordToGuess)
					return
				}
				continue
			}
		}
		guessedLetter := rune(guess[0])
		// Check if the letter has already been guessed
		if gameState.GuessedLetters[guessedLetter] {
			fmt.Println("You already guessed that letter.")
			continue
		}
		// Add the guessed letter to the map of guessed letters
		gameState.GuessedLetters[guessedLetter] = true
		if !strings.ContainsRune(gameState.WordToGuess, guessedLetter) {
			gameState.Attempts--
			fmt.Printf("Not present in the word, %d attempts remaining\n", gameState.Attempts)
			// Display the hangman only if the guess is incorrect
			hangman.DisplayHangman(gameState.Attempts, gameState.HangmanPositions)
		}
		revealedWord := game.RevealLetters(gameState.WordToGuess, gameState.GuessedLetters)
		displayedWordWithoutSpaces := strings.ReplaceAll(revealedWord, " ", "")
		// Check if the user won
		if displayedWordWithoutSpaces == gameState.WordToGuess {
			fmt.Println("\nYou won!\n\n +--^----------,--------,-----,--------^-,\n | |||||||||   `--------'     |          O\n `+---------------------------^----------|\n   `\\_,---------,---------,--------------'\n     / XXXXXX /'|       /'\n    / XXXXXX /  `\\    /'\n   / XXXXXX /`-------'\n  / XXXXXX /\n / XXXXXX /\n(________(                \n `------'")
			return
		}
	}
	fmt.Println("Haha, you lost! The word was:", gameState.WordToGuess)
}
