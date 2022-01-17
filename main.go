package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func loadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

// loadWordsMap loads all words from file and if the word has duplicate characters
func loadWordsMap(filename string) (map[string]bool, error) {
	wordList, err := loadFile(filename)
	if err != nil {
		return nil, err
	}
	result := make(map[string]bool, len(wordList))
	for _, word := range wordList {
		result[word] = hasDuplicateChar(word)
	}
	return result, nil
}

func main() {
	words, err := loadWordsMap("words.txt")
	if err != nil {
		panic("failed to load file")
	}

	fmt.Printf("total %d words found\n", len(words))
	// blacklistWords := []string{}

	tryWord := "raise"
	tryCount := 0
	for {
		fmt.Printf("Try [%s] and type the result(0 for not-exist, 1 for exist, 2 for correct, -1 for no such word) > ", tryWord)

		var input string
		fmt.Scanln(&input)
		tryCount++

		if input == "22222" {
			fmt.Printf("You win!\n")
			break
		}

		if input != "-1" {
			words = filterByLastGuess(words, tryWord, input)
		} else {
			delete(words, tryWord)
		}

		if len(words) == 0 {
			fmt.Printf("no words found\n")
			return
		}

		// now select a new word to try
		// NOTE: select word that has no duplicate characaters
		tryWord = findNewWordToTry(words)
	}
}

func hasDuplicateChar(word string) bool {
	charMap := make(map[string]bool)
	for index := range word {
		if _, ok := charMap[word[index:index+1]]; ok {
			return true
		}
		charMap[word[index:index+1]] = true
	}
	return false
}

func findNewWordToTry(words map[string]bool) string {
	// first try to find a word that has no duplicate characaters
	for word, dup := range words {
		if !dup {
			return word
		}
	}

	// since all words has duplicate characaters, return a random word
	for word := range words {
		return word
	}

	// this should never happen
	return ""
}

// filterByLastGuess returns a new slice of words that fits last guess
func filterByLastGuess(words map[string]bool, tryWorld, guessResult string) map[string]bool {
	filteredWords := make(map[string]bool, 0)
	for word, dup := range words {
		passed := true
		for i := 0; i < len(tryWorld); i++ {
			checkChar := tryWorld[i : i+1]
			checkResult := guessResult[i : i+1]
			switch checkResult {
			case "0":
				if strings.Contains(word, checkChar) {
					passed = false
					break
				}
			case "1":
				if !strings.Contains(word, checkChar) {
					passed = false
					break
				}

				if checkChar == word[i:i+1] {
					passed = false
					break
				}
			case "2":
				if checkChar != word[i:i+1] {
					passed = false
					break
				}
			}
		}

		if passed {
			filteredWords[word] = dup
		}
	}

	return filteredWords
}
