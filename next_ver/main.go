package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

var userInput string

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Welcome to VIPER Lang\n> ")

	userInput , _ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)
	
	fmt.Print("input :-"+userInput+"\n")

	interpret(userInput)
}

func interpret(input string) {

	if strings.HasPrefix(input,"run ") {
		parts := []rune(input)
		fileName := string(parts[4:])
	}
	
}