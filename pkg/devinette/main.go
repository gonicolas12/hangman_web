package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const number = 7

func guessHandler(w http.ResponseWriter, r *http.Request) {
	// Assurez-vous que nous utilisons la méthode POST pour la requête
	if r.Method != "POST" {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	// Parsez la supposition de l'utilisateur
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	guess, err := strconv.Atoi(r.FormValue("guess"))
	if err != nil {
		http.Error(w, "Invalid number", http.StatusBadRequest)
		return
	}

	// Comparez la supposition avec le nombre prédéfini et renvoyez le résultat
	if guess == number {
		fmt.Fprintf(w, "Félicitations ! Vous avez deviné juste !")
	} else {
		fmt.Fprintf(w, "Désolé, c'est incorrect. Le bon nombre est %d", number)
	}
}

func main() {
	http.HandleFunc("/guess", guessHandler)

	// Ajoutez cette ligne pour servir votre fichier HTML
	http.Handle("/", http.FileServer(http.Dir("../../web/templates")))

	fmt.Println("Le serveur fonctionne sur le port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
