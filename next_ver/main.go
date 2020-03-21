package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"io/ioutil"
	"reflect"

	//"compiler/cmd"
)
const (
	//errors
	CMD_ARG_ERR string= "\u001B[91mINVALID COMMAND\u001B[0m\n"
	METHOD_DECLARATION_ERR string = "The '{' token should not have any following token except new line feed or space"
	INSUFFICIENT_VARS_ERR string = "the number of vars on the lhs dont match with the values on the rhs"
	MAIN_MISSING_ERR string = "FATAL ERROR\nMETHOD MAIN IS MISSING"
	BUILD_FAIL_ERR string = "FATAL ERROR\nBUILD FAILED"

	//files
	DEFAULT_EXEC_FILE = "lethalityTest.vpr"
	//for method syntactic sugar
	METHOD_SYNTACTIC_SUGAR_STORAGE = "methodSugar.txt"
	SUGAR_DELIM = "|"

	//syntax
	FUNCTION_CALL string = "<<"
	NORMAL_ASSIGNMENT string = "="
	SYNTACTIC_ASSIGNMENT string = "<-"

	RETURN_STATEMENT string = "return "

	COMMENT_START string = "#"

	METHOD_DECLARATION string = "method"
	VAR_DECALRATION string = "var "
	SCOPE_DECLARATION string = "scope "

	//special syntax
	PRINTING string = "print "

	//special constants
	NULL = "null"

)
//structs
type actionFunc func(string) (bool,string)
type arith func(string,string) (string)

type CmdArgs struct {
	action actionFunc
}

type Data struct {
	value string
}

type MethodData struct {
	parameters string
	data Block
	scopeNode ScopeNode
}

type ScopeNode struct {
	parent *ScopeNode
	varSave map[string]Data
}

type VarSyntaxData struct {
	data Block
}

type Block struct {
	data []string
}

var varSyntaxSave map[string]VarSyntaxData
var varSyntaxNames []string
var methodSave map[string]*MethodData
var methodNames []string
var globalScope ScopeNode

var buldFailure bool

func main() {
	argHandler := InititializeCMD()
	args := getUserArgs()
	
	if args != "shell" {
		methodSave = make(map[string]*MethodData)
		varSyntaxSave = make(map[string]VarSyntaxData)
		globalScope = ScopeNode{nil,make(map[string]Data)}

		
		program := Interpret(args,argHandler)
		startExec(program)
	} else {
		fmt.Println("WELCOME TO VIPER LANG")
	
		for true {
			methodSave = make(map[string]*MethodData)
			varSyntaxSave = make(map[string]VarSyntaxData)
			globalScope = ScopeNode{nil,make(map[string]Data)}
			userInput := ""
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("\u001B[92m> \u001B[0m")

			userInput ,_ = reader.ReadString('\n')
			userInput = strings.Replace(userInput, "\n", "", -1)
			program := Interpret(userInput,argHandler)

			startExec(program)

			nullifyData()
		}
	}
}

func nullifyData() {
	varSyntaxSave = nil
	methodSave = nil
	globalScope = ScopeNode{}
	methodNames = nil
	varSyntaxNames = nil
	buldFailure = false
}

//initializes the commands
func InititializeCMD() [3]CmdArgs{
	var argHandler[3] CmdArgs

	argHandler[0].action = func(args string) (bool,string){
		if (strings.TrimSpace(args) == "run -d") {
			str := CommenceReading(DEFAULT_EXEC_FILE)
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

func getUserArgs() string{
	if len(os.Args) == 1 {
		return "shell"
	}

	str := ""

	for i := 1; i < len(os.Args); i++ {
		str = str+os.Args[i]+" "
	}
	return str
}

//reads the file
func CommenceReading(fileName string) string {
	data,ERR := ioutil.ReadFile(fileName)

	if ERR != nil {
		fmt.Print("\u001B[91m",ERR,"\u001B[0m\n")
		os.Exit(0)
		return ""
	}

	fmt.Println("\u001B[92mFILE OPENED\u001B[0m\n")

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

			i = globalScope.declareMethod(name,parameters,program,i)
		}
	}

	if buldFailure == true {
		fmt.Println("\u001B[91m"+BUILD_FAIL_ERR,"\u001B[0m")
	} else {
		fmt.Println("\u001B[92mFIRST ASSIGNMENT RUN \nSTATUS: COMPLETE\u001B[0m\n")
		StartExecution()
	}
}

func (caller ScopeNode)declareMethod (name string,parameters string,program []string,index int) (int){
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
	block,endIndex := getBlock(program,index,'{','}')

	node := MethodData{param,block, ScopeNode{&caller,make(map[string]Data)}}

	methodSave[name] = &node

	methodNames = append(methodNames,name)

	return endIndex
}

func getBlock(program []string,startIndex int,start rune,end rune) (Block,int){
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
	if (startIndex > endIndex) {
		fmt.Println("\u001B[91m"+METHOD_DECLARATION_ERR,"\u001B[0m\n","line:\t",startIndex+1)
		buldFailure = true
		return Block{},startIndex
	}
	snippet := make([]string,endIndex-startIndex)
	for i := startIndex; i <endIndex; i++ {
		snippet[i-startIndex] = program[i]
	}
	return Block{snippet},endIndex
}
func declarable (line string) (bool) {
	return strings.HasPrefix(line,VAR_DECALRATION)
}

func (caller ScopeNode)declare (line string,program []string,i int,check bool) {
	if strings.HasPrefix(line,VAR_DECALRATION) {
		parts := []rune(line)
		name := string(parts[len(VAR_DECALRATION):])

		if strings.Contains(name,SYNTACTIC_ASSIGNMENT) {
			temp := strings.Split(name,SYNTACTIC_ASSIGNMENT)
			name = strings.TrimSpace(temp[0])
			if testVarDeclaration(name) {
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

				fmt.Println("\u001B[91m"+INSUFFICIENT_VARS_ERR,"\u001B[0m\n","line:\t",i+1)
				buldFailure = true
			}
			if len(allValues) == 1 {
				caller.assignToAll(allVarNames,allValues[0])
			} else {
				caller.assignAll(allVarNames,allValues)
			}
		}
	}
}

func declareVarSyntax(name string,program []string,index int) ([]string){	
	block,endIndex := getBlock(program,index,'{','}')
	varSyntaxSave[name] = VarSyntaxData{block}
	varSyntaxNames = append(varSyntaxNames,name)

	program = deleteBlock(program,index,endIndex)

	return program
}

func StartExecution() {
	for i := 0; i < len(methodNames); i++ {
		param := methodSave[methodNames[i]].parameters
		data := methodSave[methodNames[i]].data.data
		scopeNode := methodSave[methodNames[i]].scopeNode

		methodSave[methodNames[i]] = &MethodData{param , Block{replaceSpecialSyntax(data)} , scopeNode }

		fmt.Println("\u001B[92mVAR SYNTAX REPLACEMENT IN METHOD [",methodNames[i],"] \nSTATUS: COMPLETE\u001B[0m\n")
	}
	err := callFunctionMAIN()
	if err != "" {
		fmt.Println("\u001B[91m"+err,"\u001B[0m\n")
	}
}

func replaceSpecialSyntax(program []string) ([]string){
	for i := 0; i < len(program); i++ {
		for j := 0; j < len(varSyntaxNames); j++ {
			name := varSyntaxNames[j]
			value := varSyntaxSave[name].data.data
			program = Insert(program,name,value)
		}
	}

	return program
}

func Insert(program []string,name string,value []string) []string{
	emptyArray := make([]string,len(value))

	for i := 0; i < len(program); i++ {
		if (strings.TrimSpace(program[i])) == name {
			program[i] = ""

			for k := 0;k<len(emptyArray);k++{
				program = append(program,emptyArray[k])
			}

			program = moveRight(program,i,len(value))

			for k:=1;k<len(value);k++ {
				program[i+k] = value[k]
			}
		}
	}

	return program
}

func moveRight (program []string,index int,times int) []string{
	output := make([]string,len(program)-times)
	for i:= 1;i<index;i++ {
		output[i] = program[i]
	}
	for i := len(program)-1;i>times;i-- {
		output = append(output,program[i])
	}


	return output
}

func callFunctionMAIN() (string){
	fmt.Println("\u001B[92mEXECUTION ENVIRONMENT SET\nCODE EXECUTION STARTING\u001B[0m\n\n")
	main := methodSave["main"]
	var data MethodData = MethodData{main.parameters,main.data,main.scopeNode}
	if len(data.data.data) == 0 {
		return MAIN_MISSING_ERR
	}
	data.runThrough(1)
	return ""
}

func (data MethodData) runThrough (index int) {
	program := data.data.data
	for i := index; i < len(program); i++ {

		fmt.Println("line:\t",i)

		if declarable(program[i]) {
			data.scopeNode.declare(program[i],program,i,true)
		}
		
		if assignable(program[i]) {
			if strings.Contains(program[i],NORMAL_ASSIGNMENT) {
				data.scopeNode.assign(program,NORMAL_ASSIGNMENT,i)
				continue
			}

			if strings.Contains(program[i],SYNTACTIC_ASSIGNMENT) {
				data.scopeNode.assign(program,SYNTACTIC_ASSIGNMENT,i)
				continue
			}
		}

		if scopeDeclaration(program[i]) {
			i = data.scopeNode.declareScope(program,i)
		}

		funcCall,methodName := functionCall(program[i])

		if  funcCall {
			callFunction(program[i],methodName)
		}

		data.scopeNode.checkSpecialFunctions(program[i])
		
		if strings.HasPrefix(program[i],RETURN_STATEMENT) {
			functionReturn(program[i])
		}
	}

	data = MethodData{}
}


func assignable(line string) (bool) {
	return (strings.Contains(line,NORMAL_ASSIGNMENT) || strings.Contains(line,SYNTACTIC_ASSIGNMENT))
}

func (caller ScopeNode) assign (program []string,assignmentType string,i int) {
	line := program[i]
	if assignmentType == SYNTACTIC_ASSIGNMENT {
		temp := strings.Split(line,SYNTACTIC_ASSIGNMENT)
		name := strings.TrimSpace(temp[0])
		if testVarDeclaration(name) {
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
			fmt.Println("\u001B[91m"+INSUFFICIENT_VARS_ERR,"\u001B[0m\n","line:\t",i+1)
		}
		if len(allValues) == 1 {
			caller.assignToAll(allVarNames,allValues[0])
		} else {
			caller.assignAll(allVarNames,allValues)
		}
	}
}

func deleteBlock (program []string,startIndex int,endIndex int) ([]string){
	for i := startIndex; i <= endIndex; i++ {
		program[i] = ""
	}

	return program
}


func (caller ScopeNode)assignToAll(varNames []string,value string) {
	data := caller.compute(value)

	for i := 0; i < len(varNames); i++ {
		varNames[i] = strings.TrimSpace(varNames[i])
		if strings.HasPrefix(varNames[i],VAR_DECALRATION) {
			varNames[i] = string([]rune(varNames[i])[len(varNames[i]):])
		}
		caller.varSave[varNames[i]] = data
	}
}
func (caller ScopeNode) assignAll(varNames []string,varValues []string) {
	for i := 0; i < len(varNames); i++ {
		varNames[i] = strings.TrimSpace(varNames[i])
		if strings.HasPrefix(varNames[i],VAR_DECALRATION) {
			varNames[i] = string([]rune(varNames[i])[len(varNames[i]):])
		}
		data := caller.compute(varValues[i])
		caller.varSave[varNames[i]] = data
	}
}

func testVarDeclaration(name string) bool{
	return true
}

func functionCall(args string) (bool,string){
	for i := 0; i < len(methodNames); i++ {
		if (strings.HasPrefix(args,methodNames[i])) {
			return true,methodNames[i]
		}
	}
	return false,""
}

func callFunction (method string,name string) {
	method = removeSyntacticSugar(strings.TrimSpace(method))
	param := string([]rune(method)[len(name)+1:])

	assignmentSting := methodSave[name].parameters+" = "+param
	methodSave[name].data.data[0] = assignmentSting

	methodSave[name].runThrough(0)
}

func removeSyntacticSugar(method string) string{
	data,ERR := ioutil.ReadFile(METHOD_SYNTACTIC_SUGAR_STORAGE)

	if ERR != nil {
		fmt.Print("\u001B[91m",ERR,"\u001B[0m\n")
		os.Exit(0)
		return ""
	}

	allSugar := strings.Split(string(data),"\n")

	methArray := strings.Split(method," ")

	returnArray := make([]string,len(methArray))

	for i := 0; i < len(allSugar); i++ {
		sugar := strings.Split(allSugar[i],SUGAR_DELIM)

		for j := 0; j < len(returnArray); j++ {
			for k:=0;k<len(sugar);k++ { 

				returnArray[j] = methArray[j]
				
				if strings.TrimSpace(methArray[j]) == sugar[k] {
					returnArray[j] = sugar[1]
					break
				}
			}
		}
	}

	returnString :=""

	for i := 0; i < len(returnArray); i++ {
		returnString = returnString+returnArray[i]+" "
	}

	return returnString
}

func scopeDeclaration(code string) bool{
	return strings.HasPrefix(code,SCOPE_DECLARATION)
}

func (caller ScopeNode) declareScope (program []string,index int) (int){
	block,endIndex := getBlock(program,index,'{','}')
	scope := MethodData{"",block,ScopeNode{&caller,make(map[string]Data)}}
	scope.runThrough(0)
	return endIndex
}

func (caller ScopeNode) checkSpecialFunctions(line string) {
	if strings.HasPrefix(line,PRINTING) {
		caller.print(line)
	}
}

func (caller ScopeNode) print(statement string) {
	statement = strings.TrimSpace(statement)
	statement = string([]rune(statement)[len(PRINTING)+1:len(statement)-1])
	statement = caller.replaceVars(statement)
	//fmt.Println("\u001B[96m"+statement+"\u001B[0m")
	fmt.Println(statement)
}

func (caller ScopeNode) replaceVars (statement string) (string){
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
			value,found := caller.searchTreeFor(pureName)
			if found {
				statement = strings.Replace(statement,varName,value,1)
			} else {
				statement = strings.Replace(statement,varName,NULL,1)
			}
			name = make([]rune,0)
		}
	}

	return statement
}

func (caller ScopeNode)searchTreeFor(name string) (string,bool){
	found := false
	itrnext := true
	presentNode := &caller
	for !found {
		if itrnext {
			data := presentNode.varSave[name].value
			if data != "" {
				return data,true
			}
			if (reflect.DeepEqual(presentNode.parent.varSave,globalScope.varSave)){
				presentNode = presentNode.parent
				itrnext = false
				found = true
			} else {
				presentNode = presentNode.parent
			}
		} else {
			return "",false
		}
	}

	return "",false
}

func functionReturn(line string) {

}

//compute is incomplete ##########################################3
func (caller ScopeNode) compute(value string) (Data) {
	value =  strings.TrimSpace(value)
	elems := strings.Split(value," ")
	for i := 0; i < len(elems); i++ {
		if (elems[i] != "" || elems[i] != "" ){
			value,found := caller.searchTreeFor(elems[i])
			if found {
				elems[i] = value
			}
		}
	}

	value = ""

	for i := 0; i < len(elems); i++ {
		value = elems[i] + " "
	}

	fmt.Println(value)

	return Data{value}
}
