package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Word Counter!")
	for {
		fmt.Println("Please enter the text:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		wordCount := countWords(text)
		fmt.Printf("The number of words in the text is: %d\n", wordCount)

		fmt.Println("Do you want to enter another text? (y/n)")
		scanner.Scan()
		choice := scanner.Text()
		if strings.ToLower(choice) != "y" {
			break
		}
	}
}

func countWords(text string) int {
	words := strings.Split(text, " ")
	return len(words)
}
