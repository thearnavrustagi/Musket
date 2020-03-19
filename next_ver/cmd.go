package cmd

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)

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