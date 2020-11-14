package errors

import "fmt"

func noArguments(args []string) bool { //handling "no arguments" error
	if len(args) == 0 {
		return true
	}
	return false
}

func invalidCharacter(args []string) bool { //handling "non-valid character" error
	for _, val := range args {
		for i := range val {
			if val[i] < 32 || val[i] > 126 {
				return true
			}
		}
	}
	return false
}

func HandlingError(args []string) bool {
	switch {
	case noArguments(args):
		fmt.Println("You didn't give any arguments !")
		return true
	case invalidCharacter(args):
		fmt.Println("You used a non-valid character.")
		return true
	default:
		return false
	}
}
