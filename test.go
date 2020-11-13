package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:] //on récupère les arguments dans la variable args

	for i := range args {
		fmt.Println(args[i])
	}

}
