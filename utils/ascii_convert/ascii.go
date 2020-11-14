package ascii

import "fmt"

func returnCharacter(characterNumber int, ascii []string) []string { //return an array of 8 lines for each characters
	characterNumber = characterNumber + 8*characterNumber
	var line []string
	for i := characterNumber; i < characterNumber+8; i++ {
		line = append(line, ascii[i])
	}
	return line
}

func returnCharacterLine(text string, ascii []string) []string { //return an array that contain each lines of the function above
	var lineToReturn []string
	for _, val := range text { //val = character of the word
		value := int(val) - 32
		for _, line := range returnCharacter(value, ascii) {
			lineToReturn = append(lineToReturn, line)
		}
	}
	return lineToReturn
}

func PrintTextInAscii(text string, ascii []string) {
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
