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

func delChar(s string, index int) string {
	return s[:index] + s[index+1:]
}

func checkWordWithLastGuess(checkWord string, guessWord, guessResult string) bool {
	if len(checkWord) == 1 && guessResult == "2" {
		return true
	}

	passed := true
	for i := 0; i < len(guessWord); i++ {
		checkChar := guessWord[i : i+1]
		checkResult := guessResult[i : i+1]
		switch checkResult {
		case "0":
			//  检查的字符不存在有两种可能：
			// 1. 字符串中不包含该字符
			// 2. 字符串中有该字符串，但是猜测的单词中有多个该字符串，并且其他字符串已经完全匹配
			correctGuessCount := 0
			for j := 0; j < len(guessWord); j++ {
				if guessResult[j:j+1] == "2" && guessWord[j:j+1] == checkChar {
					correctGuessCount += 1
				}
			}

			if strings.Count(checkWord, checkChar) != correctGuessCount {
				passed = false
				break
			}
		case "1":
			if !strings.Contains(checkWord, checkChar) {
				passed = false
				break
			}

			if checkChar == checkWord[i:i+1] {
				passed = false
				break
			}
		case "2":
			if checkChar != checkWord[i:i+1] {
				passed = false
				break
			}
		}
	}
	return passed
}

// filterByLastGuess returns a new slice of words that fits last guess
// guess result is a string of 0, 1, 2:
// * 0 means the character does not exist
// * 1 means the character exists but not in the right position
// * 2 means the character exists and in the right position
//
// NOTE: there could be a situation where the guess word has duplicate characters, and guess result should be interpreted as the following:
// * if one of the duplicate characters is in the right postition, the other character guess result indicates the match result of all but the first character.
//   e.g. the guesss word is "plook", and the guess result is "01201", the first "o"'s guess result is "2" which means it's in the right position,
//   the second "o"'s guess result is "0" which means there is only one "o" in the word, not "o"  does not exist
func filterByLastGuess(words map[string]bool, guessWord, guessResult string) map[string]bool {
	filteredWords := make(map[string]bool, 0)
	for word, dup := range words {
		if checkWordWithLastGuess(word, guessWord, guessResult) {
			filteredWords[word] = dup
		}
	}

	return filteredWords
}
