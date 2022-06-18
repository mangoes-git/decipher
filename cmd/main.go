package main

import (
	"bufio"
	"fmt"
	"os"
)

func run(source string) {
	// initialize lexer
	// scan tokens
	// print tokens one by one
	return
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Print("couldn't read file")
	}
	run(string(bytes))
}

func runPrompt() {
	s := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		status := s.Scan()
		line := s.Text()
		if !status && s.Err() == nil {
			fmt.Println("Exiting...")
			break
		}
		fmt.Println(line)
	}
}

func main() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: decipher [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		// RunFile(args[1])
	} else {
		runPrompt()
	}
}