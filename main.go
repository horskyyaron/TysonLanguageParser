package main

import (
	"fmt"
	"os"
	"os/user"
	"tyson/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Tyson programming language!\n", user.Username)
	fmt.Printf("Tyson is the greatest dog ever!\n")
	fmt.Printf("That is why he even has a programming language named after him!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
