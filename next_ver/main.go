package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

var userInput string

//put this in another package called cmd.go
type actionFunc func(string) bool
type CmdArgs struct {
	action actionFunc
}

func InititializeCMD() [3]CmdArgs{
	var argHandler[3] CmdArgs

	defFile := "lethalityTest.vpr"

	argHandler[0].action = func(args string) bool{
		if (strings.TrimSpace(args) == "run -d") {
			CommenceReading(defFile)
			return true
		}
		return false
	}

	argHandler[1].action = func(args string) bool{

		if strings.HasPrefix(args,"run ") {
			parts := []rune(args)
			fileName := string(parts[4:])

			CommenceReading(fileName)
			return true
		}

		return false

	}

	argHandler[2].action = func(args string) bool{
		if args == "quit"||args == "exit" {
			os.Exit(0)
		}
		return false
	}

	return argHandler
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Welcome to VIPER Lang\n> ")

	userInput ,_ := reader.ReadString('\n')
	userInput = strings.Replace(userInput, "\n", "", -1)

	argHandler := InititializeCMD()
	Interpret(userInput,argHandler)
}

// upgrade with structs and higher order functions
func Interpret(input string,argHandler[3] CmdArgs) {
	for i := 0; i < len(argHandler); i++ {
		success := argHandler[i].action(input)

		if success {
			os.Exit(0)
		}
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