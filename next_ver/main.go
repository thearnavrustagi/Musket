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

	Interpret(userInput)
}

// upgrade with structs and higher order functions
func Interpret(input string) {

	if strings.HasPrefix(input,"run ") {
		parts := []rune(input)
		fileName := string(parts[4:])

		CommenceReading(fileName)
	}
}

func CommenceReading(fileName string) {
	file,ERR := os.Open(fileName)

	if ERR != nil {
		fmt.Print(ERR)
		os.Exit(0)
	}


	defer file.Close()
}