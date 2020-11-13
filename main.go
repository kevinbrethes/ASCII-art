// In this project we want to convert the passed arguments into an ascii art
// We'll use the file "standard.txt" which contain the ASCII table
// Each characters is composed of 8 lines

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanLines(path string) ([]string, error) { //put each lines of the txt file in an array

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func returnCharacter(characterNumber int, ascii []string) []string { //return an array of 8 lines for each characters
	characterNumber = characterNumber + 8*characterNumber
	var line []string
	for i := characterNumber; i < characterNumber+8; i++ {
		line = append(line, ascii[i])
	}
	return line
}

func returnCharacterLine(text []string, ascii []string) []string { //return an array that contain each lines of the function above
	var lineToReturn []string
	for _, val := range text { //val = word
		for i := range val { //val[i] = character of the word
			value := int(val[i]) - 32
			for _, line := range returnCharacter(value, ascii) {
				lineToReturn = append(lineToReturn, line)
			}
		}
	}
	return lineToReturn
}

func printTextInAscii(text []string, ascii []string) {
	characters := returnCharacterLine(text, ascii)

	for i := 0; i < 8; i++ {
		fmt.Print(characters[i])
		j := 8
		for i+j < len(characters) {
			fmt.Print(characters[i+j])
			j += 8
		}
		fmt.Println()
	}
}

func backlineSupport(textWithSpaces []string) [][]string {
	var finalText [][]string
	textToAdd := ""
	var arrayToAdd []string
	for _, sentence := range textWithSpaces {
		for i := 1; i < len(sentence)-1; i++ {
			if sentence[i] == '\\' && sentence[i+1] == 'n' {
				arrayToAdd = append(arrayToAdd, textToAdd)
				finalText = append(finalText, arrayToAdd)
				textToAdd = ""
				arrayToAdd = nil
			} else if sentence[i-1] != '\\' && sentence[i] != 'n' {
				if i == 1 {
					textToAdd += string(sentence[i-1])
				}
				textToAdd += string(sentence[i])
			}
		}
		if sentence[len(sentence)-1] != 'n' && sentence[len(sentence)-2] != '\\' {
			textToAdd += string(sentence[len(sentence)-1])
			arrayToAdd = append(arrayToAdd, textToAdd)
			finalText = append(finalText, arrayToAdd)
		}
	}

	return finalText
}

func main() {
	arguments := os.Args[1:] //take all the parameters

	if len(arguments) == 0 { //handling "no arguments" error
		fmt.Println("You didn't give any arguments !")
		return
	}

	for _, val := range arguments { //handling "non-valid character" error
		for i := range val {
			if val[i] < 32 || val[i] > 126 {
				fmt.Println("You used a non-valid character.")
				return
			}
		}
	}

	var textWithSpaces []string
	for i, val := range arguments { //add a space between each arguments

		textWithSpaces = append(textWithSpaces, val)
		if i < len(arguments)-1 {
			textWithSpaces = append(textWithSpaces, " ")
		}
	}

	lines, err := scanLines("./standard.txt") //put each lines of the txt file in the array lines
	if err != nil {
		log.Fatal(err)
		return
	}

	finalText := backlineSupport(textWithSpaces)

	for _, val := range finalText {
		printTextInAscii(val, lines)
	}
}
