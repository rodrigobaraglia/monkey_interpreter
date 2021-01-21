package main

import (
	"fmt"
	"os"
	"os/user"
	"rodrigobaraglia/monkeyimpl/repl"
)

func main() {

	fmt.Println("Wait...")
	user, err := user.Current()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to the MonkeyLang REPL!\n", user.Username)
	fmt.Println("Type any command and see what happens")
	repl.Start(os.Stdin, os.Stdout)

}
