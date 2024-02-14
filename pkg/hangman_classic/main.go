package main

import (
	"fmt"
	"hangman-classic/game"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/guess", game.GameHandler)
	http.Handle("/", http.FileServer(http.Dir("../../web/templates/")))
	fmt.Println("Le serveur fonctionne sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
