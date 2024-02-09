package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "run", "./main.go", "./words.txt", "./hangman.txt")
	cmd.Dir = "../pkg/hangman_classic"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
