package main

import (
	"fmt"
	"strings"
)

var userInput string;

func main() {
	fmt.Print("Welcome to VIPER Lang\n> ")
	fmt.Scanln(&userInput)

	interpret()
}

func interpret() {
	fmt.Print("trying")
	if strings.HasPrefix(userInput,"run") {
		fmt.Print("in if %d\n",len(userInput))
		parts := []rune(userInput)
		fmt.Print("slice %v",parts)
		fileName := string(parts[:])
		fmt.Print("flname :-"+fileName+"\n")
	}

	fmt.Print("elsr")
}