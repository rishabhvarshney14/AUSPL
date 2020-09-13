package main

import (
	"fmt"
	"auspl/repl"
	"os"
)

func main() {
	fmt.Printf("Welcome to AUSPL!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
