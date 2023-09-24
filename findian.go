// The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. The program should print “Found!” if the entered string starts with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a word: ")
	reader := bufio.NewReader(os.Stdin)
	word, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	word = strings.TrimSpace(word)
	word = strings.ToLower(word)
	wordLength := len(word)
	lastChar := word[wordLength-1]
	firstChar := word[0]
	fmt.Printf("fist char: %c   last char: %c \n", firstChar, lastChar)
	switch {
	case wordLength < 3:
		{
			fmt.Println("Not Found!")
		}
	case strings.Contains(word, "a") && lastChar == 'n' && firstChar == 'i':
		{
			fmt.Println("Found!")
		}
	default:
		{
			fmt.Println("Not Found!")
		}
	}
}
