package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"io/ioutil"

	//"compiler/cmd"
)
//put this in another file called errors or constants
const (
	CMD_ARG_ERR string = "INVALID COMMAND"
)
//.........................................

func main() {
	fmt.Println("Welcome to VIPER Lang")
	for true {
		userInput := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("> ")

		userInput ,_ = reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		argHandler := InititializeCMD()
		program := Interpret(userInput,argHandler)

		check(program)
	}
}

func check(args string) {
	if args == CMD_ARG_ERR {
		fmt.Print("\u001B[91m",CMD_ARG_ERR,"\u001B[0m\n")
	} else {
		AssignmentRun(args)
	}
}

// this should also be in a file called cmd #####################################
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

func Interpret(input string,argHandler[3] CmdArgs) string{
	for i := 0; i < len(argHandler); i++ {
		
		success,data := argHandler[i].action(input)


		if success {
			return data
		}
	}
	return CMD_ARG_ERR
}

func CommenceReading(fileName string) string {
	data,ERR := ioutil.ReadFile(fileName)

	if ERR != nil {
		fmt.Print("\u001B[91m",ERR,"\u001B[0m\n")
		return ""
	}
	program := string(data)

	return program
}
//.............................................................................

//put this code in a file called verifier#################################

type idCheck func(string) bool

type Syntax struct {
	identity idCheck
	name string
}

func AssignmentRun(program string) {
	//the splitting on new line works
	splitCode := strings.Split(program,"\n")

	for i := 0; i < len(splitCode); i++ {

	}
}
//...........................................................
