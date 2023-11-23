package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Hangman struct {
	WordToGuess        string
	HiddenWord         []rune
	DisplayedWord      string
	UserInput          string
	GameState          string
	LatestPropositions []string
	Attempts           int
}

var WORDS = []string{}

func LoadWords(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	word := ""
	for _, char := range string(content) {
		if char == '\n' {
			WORDS = append(WORDS, strings.ToUpper(word[:len(word) - 1]))
			word = ""
		} else {
			word += string(char)
		}
	}
}

func InitHangman() {
	rand.Seed(time.Now().UTC().UnixNano())

	hangman = Hangman{
		WordToGuess:        strings.ToUpper(WORDS[rand.Intn(len(WORDS))]),
		HiddenWord:         []rune{},
		Attempts:           0,
		LatestPropositions: []string{},
	}
	for range hangman.WordToGuess {
		hangman.HiddenWord = append(hangman.HiddenWord, '_')
	}
	refreshHangman()
}

func TestInput(input string) {
	if len(input) != 1 {
		return
	}

	alreadyIn := false
	for _, proposition := range hangman.LatestPropositions {
		if proposition == strings.ToUpper(input) {
			alreadyIn = true
			break
		}
	}

	if alreadyIn {
		fmt.Println("Cette lettre a déjà été proposée.")
		return
	}

	hangman.LatestPropositions = append(hangman.LatestPropositions, strings.ToUpper(input))

	if len(input) == 1 {
		success := false
		for i := 0; i < len(hangman.WordToGuess); i++ {
			if string(hangman.WordToGuess[i]) == strings.ToUpper(input) {
				success = true
				hangman.HiddenWord[i] = rune(strings.ToUpper(input)[0])
			}
		}
		if !success {
			hangman.Attempts += 1
		}
	} else {
		if strings.ToUpper(input) == hangman.WordToGuess {
			hangman.HiddenWord = []rune(hangman.WordToGuess)
		} else {
			hangman.Attempts += 2
			if hangman.Attempts > 10 {
				hangman.Attempts = 10
			}
		}
	}

	refreshHangman()
}

func Seq(count int) []int {
	result := make([]int, count)
	for i := range result {
		result[i] = i + 1
	}
	return result
}

func refreshHangman() {
	hangman.DisplayedWord = string(hangman.HiddenWord)
	if string(hangman.HiddenWord) == hangman.WordToGuess {
		hangman.DisplayedWord = "Congratulation ! You WIN !!!"
	} else if hangman.Attempts >= 10 {
		hangman.DisplayedWord = "Congratulation ! You're a looser..."
	}
}
