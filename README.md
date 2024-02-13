# Hangman Web

## Overview
Welcome to the Hangman game ! This classic word-guessing game challenges players to guess a word by suggesting letters within a certain number of guesses.

## Installation
1. Clone the repository.
2. Ensure Go is installed on your machine.
3. Execute the program with `go run main.go` in the `cmd` directory

## How to Play
- Click on `Jouer` to begin a game.
- Select a difficulty
- You'll see the word represented by underscores, indicating the number of letters.
- Guess a letter at a time, or attempt to guess the whole word.
- Each incorrect guess brings you closer to losing the game.
- The game ends when you either guess the word correctly or run out of attempts.

## Features
- Customizable word list and hangman stages.
- Button to put the music on and off.
- Other buttons which works.
- Different pages when you lose or when you win.
- Different musics free to use (I think).

## Difficulties
- Easy : the number of letters that are revealed at the beginning of the game is the length of the word divided by 2 minus 1
- Average : the number of letters that are revealed at the beginning of the game is the length of the word divided by 2 minus 2
- Hard : only one letter is reavealed at the beginning of the game
- Extreme : there are no letters reavealed at the beginning of the game