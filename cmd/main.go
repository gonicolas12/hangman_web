package main

// NOT finished

import (
	"fmt"
	"log"

	"hangman_web/pkg/hangman_classic"

	"../web"
)

func main() {
	// Create a new instance of the Hangman game
	game := hangman_classic.NewGame()

	// Create a new instance of the Hangman web server
	server := web.NewServer(game)

	// Start the web server
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hangman game started. Open your browser and visit http://localhost:8080 to play!")
}
