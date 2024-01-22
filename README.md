# Hangman Web

## Overview
Welcome to the Hangman game ! This classic word-guessing game challenges players to guess a word by suggesting letters within a certain number of guesses.

## Installation
1. Clone the repository.
2. Ensure Go is installed on your machine.
3. Execute the program with `go run main.go words.txt hangman.txt`

## How to Play
- Run the program and a word will be chosen randomly.
- Select a difficulty
- You'll see the word represented by underscores, indicating the number of letters.
- Guess a letter at a time, or attempt to guess the whole word.
- Each incorrect guess brings you closer to losing the game.
- The game ends when you either guess the word correctly or run out of attempts.

## Features
- ASCII art representation of the hangman and welcome message.
- Ability to guess single letters or the entire word.
- Save and resume game functionality (start-and-stop)
- Customizable word list and hangman stages.

## Difficulties
- Easy : the number of letters that are revealed at the beginning of the game is the length of the word divided by 2 minus 1
- Average : the number of letters that are revealed at the beginning of the game is the length of the word divided by 2 minus 2
- Hard : only one letter is reavealed at the beginning of the game
- Extreme : there are no letters reavealed at the beginning of the game