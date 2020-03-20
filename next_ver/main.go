package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"io/ioutil"
	"strconv"

	//"compiler/cmd"
)
const (
	//errors
	CMD_ARG_ERR string= "\u001B[40m\u001B[91mINVALID COMMAND\u001B[0m\n"
	METHOD_DECLARATION_ERR string = "The '{' token should not have any following token except new line feed or space"
	INSUFFICIENT_VARS_ERR string = "the number of vars on the lhs dont match with the values on the rhs"
	ALREADY_DECLARED_ERR string = "the variables have already been declared please remove the \"var\" specifier"
	MAIN_MISSING_ERR string = "FATAL ERROR\nMETHOD MAIN IS MISSING"
	BUILD_FAIL_ERR string = "FATAL ERROR\nBUILD FAILED"

	//syntax
	FUNCTION_PARAM string = "<-"
	NORMAL_ASSIGNMENT string = "="
	SYNTACTIC_ASSIGNMENT string = "<-"

	COMMENT_START string = "#"

	METHOD_DECLARATION string = "method"
	VAR_DECALRATION string = "var "

	//special syntax
	PRINTING string = "print "

)
//structs
type actionFunc func(string) (bool,string)
type arith func(string,string) (string)

func StrToInt(str string) (int, error) {
    nonFractionalPart := strings.Split(str, ".")
    return strconv.Atoi(nonFractionalPart[0])
}

type CmdArgs struct {
	action actionFunc
}

type ArithExpression struct {
	representation string
	action arith
}

type Node struct {
	value string
	childAction string
}

type AssignmentLinkedList struct {
	child *Node
}

type Data struct {
	value string
}

type MethodData struct {
	parameters string
	data Block
}

type VarSyntaxData struct {
	data Block
}

type Block struct {
	data []string
}

//global variables
var headNode *Node

var varSave map[string]Data
var varSyntaxSave map[string]VarSyntaxData
var varSyntaxNames []string
var methodSave map[string]MethodData

var buldFailure bool

func main() {
	fmt.Println("Welcome to VIPER Lang")
	argHandler := InititializeCMD()
	
	for true {
		methodSave = make(map[string]MethodData)
		varSave = make(map[string]Data)
		varSyntaxSave = make(map[string]VarSyntaxData)
		userInput := ""
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("\u001B[92m> \u001B[0m")

		userInput ,_ = reader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)
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
		fmt.Print("\u001B[40m\u001B[91m",ERR,"\u001B[0m\n")
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
		fmt.Print(CMD_ARG_ERR)
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
		program[i] = strings.TrimSpace(program[i])

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

	if buldFailure == true {
		fmt.Println("\u001B[40m\u001B[91m",BUILD_FAIL_ERR,"\u001B[0m")
	} else {
		StartExecution(program)
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
	block := getBlock(program,index,'{','}')

	node := MethodData{param,block}

	methodSave[name] = node
}

func getBlock(program []string,startIndex int,start rune,end rune) (Block){
	num_of_nested_blocks := 0
	endIndex := startIndex

	for i := startIndex; i < len(program); i++ {
		program[i] = strings.TrimSpace(program[i])
		elems := []rune(program[i])

		for j := 0; j < len(elems); j++ {
			
			if elems[j] == '#' {
				break
			} 
			
			if elems[j] == start {
				num_of_nested_blocks++
				j = j+1
				continue
			}
			
			if elems[j] == end {
				num_of_nested_blocks--
			}
		}

		if num_of_nested_blocks == 0 {
			endIndex = i
			break
		}
	}
	if (startIndex == endIndex) {
		fmt.Println("\u001B[40m\u001B[91m",METHOD_DECLARATION_ERR,"\u001B[0m\n","line:\t",startIndex+1)
		buldFailure = true
		return Block{}
	}
	snippet := program[startIndex+1:endIndex]
	return Block{snippet}
}


func StartExecution(program []string) {
	err := callFunctionMAIN()
	if err != "" {
		fmt.Println("\u001B[40m\u001B[91m",err,"\u001B[0m\n")
	}
}

func callFunctionMAIN() (string){
	data := methodSave["main"]
	if len(data.data.data) == 0 {
		return MAIN_MISSING_ERR
	}
	data.runThrough()
	return ""
}

func (data MethodData) runThrough () {
	program := data.data.data
	for i := 0; i < len(program); i++ {
		if declarable(program[i]) {
			declare(program[i],program,i)
			continue
		}
		if assignable(program[i]) {
			if strings.Contains(program[i],NORMAL_ASSIGNMENT) {
				assign(program,NORMAL_ASSIGNMENT,i)
				continue
			}

			if strings.Contains(program[i],SYNTACTIC_ASSIGNMENT) {
				assign(program,SYNTACTIC_ASSIGNMENT,i)
				continue
			}
		}

		checkSpecialSyntax(program[i])
	}
}

func declarable (line string) (bool) {
	return strings.HasPrefix(line,VAR_DECALRATION)
}

func declare (line string,program []string,i int) {
	if strings.HasPrefix(line,VAR_DECALRATION) {
		parts := []rune(line)
		name := string(parts[len(VAR_DECALRATION):])

		if strings.Contains(name,SYNTACTIC_ASSIGNMENT) {
			temp := strings.Split(name,SYNTACTIC_ASSIGNMENT)
			name = strings.TrimSpace(temp[0])
			if testVarDeclaration(name) {
				fmt.Println("\u001B[40m\u001B[91m",ALREADY_DECLARED_ERR,"\u001B[0m\nline:\t",i+1)
				buldFailure = true
			}
			declareVarSyntax(name,program,i)
			return
		}

		if strings.Contains(name,NORMAL_ASSIGNMENT) {
			temp := strings.Split(name,NORMAL_ASSIGNMENT)
			allVarNames := strings.Split(temp[0],",")
			allValues := strings.Split(temp[1],",")

			if (len(allVarNames) != len(allValues)) && (len(allValues) != 1){

				fmt.Println("\u001B[40m\u001B[91m",INSUFFICIENT_VARS_ERR,"\u001B[0m\n","line:\t",i+1)
				buldFailure = true
			}
			if len(allValues) == 1 {
				assignToAll(allVarNames,allValues[0],i,true)
			} else {
				assignAll(allVarNames,allValues,i,true)
			}
		}
	}
}

func assignable(line string) (bool) {
	return (strings.Contains(line,NORMAL_ASSIGNMENT) || strings.Contains(line,SYNTACTIC_ASSIGNMENT))
}

func assign (program []string,assignmentType string,i int) {
	line := program[i]
	if assignmentType == SYNTACTIC_ASSIGNMENT {
		temp := strings.Split(line,SYNTACTIC_ASSIGNMENT)
		name := strings.TrimSpace(temp[0])
		if testVarDeclaration(name) {
			fmt.Println("\u001B[40m\u001B[91m",ALREADY_DECLARED_ERR,"\u001B[0m\nline:\t",i+1)
			buldFailure = true
		}

		declareVarSyntax(name,program,i)
		return
	}

	if assignmentType == NORMAL_ASSIGNMENT {
		temp := strings.Split(line,NORMAL_ASSIGNMENT)
		allVarNames := strings.Split(temp[0],",")
		allValues := strings.Split(temp[1],",")

		if (len(allVarNames) != len(allValues)) && (len(allValues) != 1){
			fmt.Println("\u001B[40m\u001B[91m",INSUFFICIENT_VARS_ERR,"\u001B[0m\n","line:\t",i+1)
		}
		if len(allValues) == 1 {
			assignToAll(allVarNames,allValues[0],i,false)
		} else {
			assignAll(allVarNames,allValues,i,false)
		}
	}
}
func declareVarSyntax(name string,program []string,index int) {	
	block := getBlock(program,index,'{','}')
	varSyntaxSave[name] = VarSyntaxData{block}
	varSyntaxNames = append(varSyntaxNames,name)
}


func assignToAll(varNames []string,value string,index int,check bool) {
	data := compute(value)
	for i := 0; i < len(varNames); i++ {
		varNames[i] = strings.TrimSpace(varNames[i])
		if strings.HasPrefix(varNames[i],VAR_DECALRATION) {
			fmt.Println(varNames[i])
			varNames[i] = string([]rune(varNames[i])[len(varNames[i]):])
		}
		if testVarDeclaration(varNames[i]) && check{
			fmt.Println("\u001B[40m\u001B[91m",ALREADY_DECLARED_ERR,"\u001B[0m\nline:\t",index+1)
		}
		varSave[varNames[i]] = data
	}
}
func assignAll(varNames []string,varValues []string,index int,check bool) {
	for i := 0; i < len(varNames); i++ {
		varNames[i] = strings.TrimSpace(varNames[i])
		if strings.HasPrefix(varNames[i],VAR_DECALRATION) {
			fmt.Println(varNames[i])
			varNames[i] = string([]rune(varNames[i])[len(varNames[i]):])
		}
		if testVarDeclaration(varNames[i]) && check{
			fmt.Println("\u001B[40m\u001B[91m",ALREADY_DECLARED_ERR,"\u001B[0m\nline:\t",index+1)
		}
		data := compute(varValues[i])
		varSave[varNames[i]] = data
	}
}

func testVarDeclaration(name string) bool{
	t1 := varSyntaxSave[name].data.data
	t2 := string(varSave[name].value)

	if (len(t1) != 0 || t2 != ""){
		return true
	}

	return false
}

func checkSpecialSyntax(line string) {
	if strings.HasPrefix(line,PRINTING) {
		print(line)
	}
}

func print(statement string) {
	statement = strings.TrimSpace(statement)
	statement = string([]rune(statement)[len(PRINTING)+1:len(statement)-1])
	statement = replaceVars(statement)
	fmt.Println(statement)
}

func replaceVars (statement string) (string){
	chars := []rune(" "+statement)
	var name []rune
	var appendable bool
	var startIndex,endIndex int

	for i := 0; i < len(chars); i++ {
		if appendable {
			name = append(name,chars[i])
		}
		if chars[i] == '{'{
			appendable = true
			startIndex = i
		}

		if chars[i] == '}'  {
			appendable = false
			endIndex = i+1
			varName := string(chars[startIndex:endIndex])
			pureName := string([]rune(varName)[1:len(varName)-1])
			value := varSave[pureName]
			statement = strings.Replace(statement,varName,value.value,1)
			name = make([]rune,0)
		}
	}
	return statement
}

//compute is incomplete ##########################################3
func compute(value string) (Data) {
	value =  strings.TrimSpace(value)
	elems := strings.Split(value," ")
	if len(elems) == 1 || strings.HasPrefix(value,"\"") && strings.HasSuffix(value,"\""){
		return Data{value}
	}
	
	return Data{}
}
