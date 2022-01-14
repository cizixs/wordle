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

func main() {
	words, err := loadFile("words.txt")
	if err != nil {
		panic("failed to load file")
	}

	fmt.Printf("total %d words found\n", len(words))

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
			words = words[1:]
		}

		if len(words) == 0 || tryCount >= 5 {
			fmt.Printf("no words found\n")
			return
		}

		// now select a new word to try
		// NOTE: select word that has no duplicate characaters
		tryWord = findNewWordToTry(words)
	}
}

func findNewWordToTry(words []string) string {
	// try to find a word that has no duplicate characaters
	for _, word := range words {
		charMap := make(map[string]bool)
		hasDuplicate := false
		for index := range word {
			if _, ok := charMap[word[index:index+1]]; ok {
				hasDuplicate = true
				break
			}
			charMap[word[index:index+1]] = true
		}
		if !hasDuplicate {
			return word
		}
	}
	return words[0]
}

// filterByLastGuess returns a new slice of words that fits last guess
func filterByLastGuess(words []string, tryWorld, guessResult string) []string {
	filteredWords := make([]string, 0)
	for _, word := range words {
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
			filteredWords = append(filteredWords, word)
		}
	}

	return filteredWords
}
