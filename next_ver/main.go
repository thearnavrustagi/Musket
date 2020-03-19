package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"io/ioutil"

	//"compiler/cmd"
)
const (
	//errors
	CMD_ARG_ERR string= "INVALID COMMAND"

	//syntax
	FUNCTION_CALL string = "<-"
	NORMAL_ASSIGNMENT string = "="
	SYNTACTIC_ASSIGNMENT string = "<-"

	COMMENT_START string = "#" 

	METHOD_DECLARATION string = "method "

	//TYPES
	//varTypes
)

//structs
type actionFunc func(string) (bool,string)

type CmdArgs struct {
	action actionFunc
}

type Data struct {
	dataType string
	value string
	syntacticType string
}

type AssignmentLinkedList struct {
	child *Node
}

type Node struct {
	data Block
	name string
}

type methodNode struct {
	parameters string
	data Block
}

type Block struct {
	data string
}

//global variables
var headNode *Node
var dataSave map[string]Data
var methodSave map[string]methodNode

func main() {
	fmt.Println("Welcome to VIPER Lang")
	methodSave = make(map[string]methodNode)
	dataSave = make(map[string]Data)

	for true {
		userInput := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("> ")

		userInput ,_ = reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		argHandler := InititializeCMD()
		program := Interpret(userInput,argHandler)

		startExec(program)
	}
}

//initializes the commands
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

//reads the file
func CommenceReading(fileName string) string {
	data,ERR := ioutil.ReadFile(fileName)

	if ERR != nil {
		fmt.Print("\u001B[91m",ERR,"\u001B[0m\n")
		return ""
	}
	program := string(data)

	return program
}

//interprets whatever commans is given
func Interpret(input string,argHandler[3] CmdArgs) string{
	for i := 0; i < len(argHandler); i++ {
		
		success,data := argHandler[i].action(input)


		if success {
			return data
		}
	}
	return CMD_ARG_ERR
}

//
func startExec(args string) {
	if args == CMD_ARG_ERR {
		fmt.Print("\u001B[91m",CMD_ARG_ERR,"\u001B[0m\n")
	} else {
		AssignmentRun(args)
	}
}

func AssignmentRun(program string) {
	//the splitting on new line works
	splitCode := strings.Split(program,"\n")

	StaticallyInitialize(splitCode)
}

func StaticallyInitialize(program []string) {
	for i := 0; i < len(program); i++ {
		if strings.HasPrefix(program[i],COMMENT_START) {
			program[i] = ""
			continue
		}

		if strings.HasPrefix(program[i],METHOD_DECLARATION) {

			parts := []rune(program[i])
			name := string(parts[len(METHOD_DECLARATION):])
			temp := strings.Split(name,NORMAL_ASSIGNMENT)
			name = strings.TrimSpace(temp[0])
			parameters := strings.TrimSpace(temp[1])

			declareMethod(name,parameters,program,i)
		}
	}
}

func declareMethod (name string,parameters string,program []string,index int) {
	elems := []rune(strings.TrimSpace(parameters))

	var startIndex,endIndex int

	for i := 0; i < len(elems); i++ {
		if elems[i] == '(' {
			startIndex = i
			continue
		}

		if elems[i] == ')' {
			endIndex = i
			break
		}
	}

	param := string(elems[startIndex+1:endIndex])
	block := getBlock(program,index)

	node := methodNode{param,block}

	methodSave[name] = node
}

func getBlock(program []string,index int) (Block){
	return Block{}
}
//...........................................................
