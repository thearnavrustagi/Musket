package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"io/ioutil"
)

//put this in another package called cmd.go till //...
type actionFunc func(string) (bool,string)
type CmdArgs struct {
	action actionFunc
}

func InititializeCMD() [3]CmdArgs{
	var argHandler[3] CmdArgs

	defFile := "lethalityTest.vpr"

	argHandler[0].action = func(args string) (bool,string){
		if (strings.TrimSpace(args) == "run -d") {
			str := CommenceReading(defFile)
			return true,str
		}
		return false,""
	}

	argHandler[1].action = func(args string) (bool,string){

		if strings.HasPrefix(args,"run ") {
			parts := []rune(args)
			fileName := string(parts[4:])

			str := CommenceReading(fileName)
			return true,str
		}

		return false,""

	}

	argHandler[2].action = func(args string) (bool,string){
		if args == "quit"||args == "exit" {
			os.Exit(0)
		}
		return false,""
	}

	return argHandler
}
//.....

func main() {
	for true {
		userInput := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Welcome to VIPER Lang\n> ")

		userInput ,_ = reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		argHandler := InititializeCMD()
		program := Interpret(userInput,argHandler)

		check(program)
	}
}

func check(args string) {
	if strings.hasSuffix(".vpr") {
		AssignmentRun(args)
	}

	else{
		fmt.Print(args,"\n")
	}
}

// this should also be in a file called cmd
func Interpret(input string,argHandler[3] CmdArgs) string{
	for i := 0; i < len(argHandler); i++ {
		
		success,data := argHandler[i].action(input)


		if success {
			return data
		}
	}
	return "Unknown command"
}

func CommenceReading(fileName string) string {
	data,ERR := ioutil.ReadFile(fileName)

	if ERR != nil {
		fmt.Print(ERR)
		os.Exit(0)
	}
	program := string(data)

	return program
}
//....

func AssignmentRun(program string) {
	
}