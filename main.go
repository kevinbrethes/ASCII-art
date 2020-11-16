// In this project we want to convert the passed arguments into an ascii art
// We'll use the file "standard.txt" which contain the ASCII table
// Each characters is composed of 8 lines

package main

import (
	"bufio"
	"log"
	"os"

	ascii "../ASCII-ART/utils/ascii_convert"
	backline "../ASCII-ART/utils/backlineSupport"
	errors "../ASCII-ART/utils/errors"
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

func main() {
	if errors.HandlingError(os.Args[1:]) { //stop the program if an error occured
		return
	}

	text := os.Args[1]

	if text == "\\n" { //avoid the backline error
		text = " "
	}

	lines, err := scanLines("./standard.txt") //put each lines of the txt file in the array "lines"
	if err != nil {
		log.Fatal(err)
		return
	}

	textArray := backline.BacklineSupport(text)

	for _, words := range textArray {
		ascii.PrintTextInAscii(words, lines)
	}
}
