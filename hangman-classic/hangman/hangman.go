package hangman

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Reads the hangman positions from the given file and returns them as a slice of strings

func ReadHangmanPositionsFromFile(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	positions := strings.Split(string(content), "\n\n")
	return positions, nil
}

// Reverse the given slice of strings

func ReverseHangmanPositions(positions []string) {
	for i, j := 0, len(positions)-1; i < j; i, j = i+1, j-1 {
		positions[i], positions[j] = positions[j], positions[i]
	}
}

// Display the hangman corresponding to the given number of attempt

func DisplayHangman(attempts int, positions []string) {
	if attempts >= len(positions) {
		attempts = len(positions) - 1
	}
	fmt.Println(positions[attempts])
}
