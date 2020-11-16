package backline

func BacklineErrorFix(text string) string { //fix the error when the args start with a "\n"
	if text[0] == '\\' && text[1] == 'n' {
		textFixed := " "
		for i := 2; i < len(text); i++ {
			textFixed += string(text[i])
		}
		return textFixed
	}
	return text
}

func BacklineSupport(text string) []string { //seperate the text into an array if there is a "\n"
	var seperateWords []string

	if len(text) == 1 { //avoid the index error
		seperateWords = append(seperateWords, text)
		return seperateWords
	}

	if text == "\\n" { //avoid the backline error
		seperateWords = append(seperateWords, " ")
		return seperateWords
	}

	words := ""

	for i := 0; i < len(text)-1; i++ {
		if text[i] != '\\' || text[i+1] != 'n' {
			words += string(text[i])
		} else {
			seperateWords = append(seperateWords, words)
			words = ""
			i++
		}
	}

	if text[len(text)-2] != '\\' || text[len(text)-1] != 'n' {
		words += string(text[len(text)-1])
		seperateWords = append(seperateWords, words)
	}

	return seperateWords
}
