package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	"io/ioutil"
	"reflect"
	"unicode"
	"strconv"

	//"compiler/cmd"
)
const (
	//errors
	CMD_ARG_ERR string= "\u001B[91mINVALID COMMAND\u001B[0m\n"
	METHOD_DECLARATION_ERR string = "The '{' token should not have any following token except new line feed or space"
	INSUFFICIENT_VARS_ERR string = "the number of vars on the lhs dont match with the values on the rhs"
	INSUFFICIENT_ARGS_FOR_INPUT_ERR string = "input statement should have 2 arguments (var,\"line\")\nthe line will br printed and \nvar will be assigned the value to"
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

	METHOD_DECLARATION string = "func "
	VAR_DECALRATION string = "var "
	SCOPE_DECLARATION string = "scope "

	//special syntax
	PRINTING string = "print "
	PURE_PRINT string = "pure_print "
	INPUT string = "input "

	//special constants
	NULL = "null"

	//data types
	NUMBER        = "|int|"
	STRING        = "|str|"
	BOOLEAN       = "|charge|"
	DOUBLE        = "|double|"
	SYSTEM_STRING = "|sys-str|"
)
//structs
type actionFunc func(string) (bool,string)
type oprAction func(Data,Data) (string,bool)

type CmdArgs struct {
	action actionFunc
}

type Data struct {
	value string
	Type string
	stringRep string
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

type Operators struct {
	rep rune
	action oprAction
}

type DataNode struct {
	data Data
	pointer NodePointer
	doesNotExist bool
}

type NodePointer struct {
	parent *DataNode
	action rune
}


var varSyntaxSave map[string]VarSyntaxData
var varSyntaxNames []string
var methodSave map[string]*MethodData
var methodNames []string
var globalScope ScopeNode
//operators
var operators map[string]Operators
var oprList []Operators

var buldFailure bool

func main() {
	argHandler := InititializeCMD()
	InititializeOperators()
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

func PrintBug(str string) {
	fmt.Println(str)
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

/*
for personal reference

PLUS
MINUS
MULT
DIVIDE
PERCENT
COMMA
*/

func InititializeOperators() {
	operators = make(map[string]Operators)
	/*pass := Operators{' ',
	func(arg1,arg2 Data) (string,bool){
		fmt.Println(arg1,"<arg1\targ2>",arg2)
		return arg1.value,true
	}}*/

	PLUS := Operators{'+',
	func(arg1 Data,arg2 Data)(string,bool){
		if arg1.Type == NUMBER {
			if arg2.Type == NUMBER {
				num1,_ := strconv.Atoi(arg1.value)
				num2,_ := strconv.Atoi(arg2.value)
				return strconv.Itoa((num1+num2)),true
			} else if arg2.Type == DOUBLE {
				dec := arg1.value+".0"
				dec1,_ := strconv.ParseFloat(dec,64)
				dec2,_ := strconv.ParseFloat(arg2.value,64)
				return strconv.FormatFloat(dec1+dec2,'f',-1,64),true
			} else if arg2.Type == BOOLEAN {
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
			} else {
				return arg1.value+arg2.value,true
			}
		} else if arg1.Type == DOUBLE {
			if arg2.Type == DOUBLE {
				num1,_ := strconv.ParseFloat(arg1.value,64)
				num2,_ := strconv.ParseFloat(arg2.value,64)
				return strconv.FormatFloat(num1+num2,'f',-1,64),true
			} else if arg2.Type == NUMBER {
				dec := arg2.value+".0"
				dec1,_ := strconv.ParseFloat(arg1.value,64)
				dec2,_ := strconv.ParseFloat(dec,64)
				return strconv.FormatFloat(dec1+dec2,'f',-1,64),true
			} else if arg2.Type == BOOLEAN {
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
			} else {
				return arg1.value+arg2.value,true
			}
		} else if arg1.Type == BOOLEAN {

			if arg2.Type == NUMBER {
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
			} else if arg2.Type == DOUBLE {
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
			} else if arg2.Type == BOOLEAN {
				val1,_ := strconv.ParseBool(arg1.value)
				val2,_ := strconv.ParseBool(arg2.value)
				return strconv.FormatBool(val1&&val2),true
			} else {
				return arg1.value+arg1.value,true
			}
		} else {
			return arg1.value+arg1.value,true
		}
		return NULL,true
	}}
	operators[string(PLUS.rep)] = PLUS

	oprList = append(oprList,PLUS)
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
			meth := string(parts[len(METHOD_DECLARATION):])
			name,parameters := getNameAndParam(meth) 

			endIndex := globalScope.declareMethod(name,parameters,program,i)

			program = deleteBlock(program,i,endIndex)
		}
		if declarable(program[i]) {
			globalScope.declare(program[i],program,i,true)
		}
	}

//	if buldFailure == true {
//		fmt.Println("\u001B[91m"+BUILD_FAIL_ERR,"\u001B[0m")
//	} else {
		fmt.Println("\u001B[92mFIRST ASSIGNMENT RUN \nSTATUS: COMPLETE\u001B[0m\n")
		StartExecution(program)
//	}
}

func getNameAndParam(meth string) (string,string){
	meth = strings.TrimSpace(meth)
	parts := []rune(meth)
	if !strings.HasSuffix(meth,"{") {
		parts = append(parts,' ')
	}
	parts[len(parts) - 1] = ' '
	for i := 0; i < len(parts); i++ {
		if parts[i] == '('{
			name := string(parts[:i])
			param := string(parts[i:])	
			return strings.TrimSpace(name),param
		}
	}

	return strings.TrimSpace(string(parts)),""
}

func (caller ScopeNode)declareMethod (name string,parameters string,program []string,index int) (int){
	elems := []rune(strings.TrimSpace(parameters))

	var startIndex,endIndex int
	var hasParam bool

	for i := 0; i < len(elems); i++ {
		if elems[i] == '(' {
			hasParam = true
			startIndex = i
			continue
		}

		if elems[i] == ')' {
			endIndex = i
			break
		}
	}
	param := ""

	if hasParam{
		param = string(elems[startIndex+1:endIndex])
	}

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

func StartExecution(program []string) {
	for i := 0; i < len(methodNames); i++ {
		param := methodSave[methodNames[i]].parameters
		data := methodSave[methodNames[i]].data.data
		scopeNode := methodSave[methodNames[i]].scopeNode

		methodSave[methodNames[i]] = &MethodData{param , Block{replaceSpecialSyntax(data)} , scopeNode }

		fmt.Println("\u001B[92mVAR SYNTAX REPLACEMENT IN METHOD [",methodNames[i],"] \nSTATUS: COMPLETE\u001B[0m\n")
	}
	err := callFunctionMAIN()
	if err != "" {
		autoCreateFunctionMAIN(program)
	}
}

func replaceSpecialSyntax(program []string) ([]string){
	for j := 0; j < len(varSyntaxNames); j++ {
		name := varSyntaxNames[j]
		value := varSyntaxSave[name].data.data
		
		program = Insert(program,name,value)
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
	if methodSave["main"] == nil {
		return MAIN_MISSING_ERR
	}
	data := MethodData{main.parameters,main.data,main.scopeNode}
	data.runThrough(0)
	return ""
}

func autoCreateFunctionMAIN (program []string) {
	main := MethodData{"",Block{program},ScopeNode{&globalScope,make(map[string]Data)}}
	main.data.data = replaceSpecialSyntax(main.data.data)
	methodSave["main"] = &main
	callFunctionMAIN()
}

func (data MethodData) runThrough (index int) {
	program := data.data.data
	for i := index; i < len(program); i++ {

		program[i] = strings.TrimSpace(program[i])

		done := data.scopeNode.checkSpecialFunctions(program[i])
		if done {
			continue
		}


		if declarable(program[i]) {
			data.scopeNode.declare(program[i],program,i,true)
			continue
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
			continue
		}

		if strings.HasPrefix(program[i],"{") && strings.HasSuffix(program[i],"}") {
			str := string([]rune(program[i])[1:len(program[i])-1])
			val,found := data.scopeNode.searchTreeFor(str)
			value := val.value
			if found{
				program[i] = value
				i = i-1
				continue
			}
		}

		funcCall,methodName := functionCall(program[i])

		if  funcCall {
			callFunction(program[i],methodName)
			continue
		}

		
		if strings.HasPrefix(program[i],RETURN_STATEMENT) {
			functionReturn(program[i])
			continue
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

func (caller ScopeNode) checkSpecialFunctions(line string) (bool){
	if strings.HasPrefix(line,PRINTING) {
		caller.print(line)
		return true
	}

	if strings.HasPrefix(line,PURE_PRINT) {
		caller.pure_print(line)
		return true
	}

	if strings.HasPrefix(line,INPUT) {
		caller.takeInput(line)
	}
	return false
}

func (caller ScopeNode) print(statement string) {
	statement = strings.TrimSpace(statement)
	statement = string([]rune(statement)[len(PRINTING)+1:len(statement)-1])
	statement = caller.replaceVars(statement)

	fmt.Println("\u001B[96m"+statement+"\u001B[0m")
}

func (caller ScopeNode) pure_print (statement string) {
	statement = strings.TrimSpace(statement)
	statement = string([]rune(statement)[len(PURE_PRINT)+1:len(statement)-1])
	statement = caller.replaceVars(statement)

	fmt.Print("\u001B[96m"+statement+"\u001B[0m")
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
				statement = strings.Replace(statement,varName,value.stringRep,1)
			} else {
				statement = strings.Replace(statement,varName,NULL,1)
			}
			name = make([]rune,0)
		}
	}

	return statement
}

func (caller ScopeNode)searchTreeFor(name string) (Data,bool){
	found := false
	itrnext := true
	presentNode := &caller
	for !found {
		if itrnext {
			data := presentNode.varSave[name].stringRep
			if data != "" {
				return presentNode.varSave[name],true
			}
			if (reflect.DeepEqual(presentNode.parent.varSave,globalScope.varSave)){
				presentNode = presentNode.parent
				itrnext = false
				found = true
			} else {
				presentNode = presentNode.parent
			}
		} else {
			return Data{},false
		}
	}

	return Data{},false
}

func (caller ScopeNode) takeInput (line string) {
	line = strings.TrimSpace(line)
	line = string([]rune(line)[len(INPUT):len(line)-1])

	parts := strings.Split(line,",")

	if len(parts) < 2 || len(parts) > 2 {
		fmt.Println("\u001B[91m"+INSUFFICIENT_ARGS_FOR_INPUT_ERR,"\u001B[0m\n")
		os.Exit(0)
	}

	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = string([]rune(parts[1])[1:])

	caller.pure_print(PURE_PRINT+" "+parts[1]+"\"")
	reader := bufio.NewReader(os.Stdin)
	str,_ := reader.ReadString('\n')
	nstr := "\""+str+"\""

	caller.varSave[parts[0]] = Data{nstr,STRING,str}
}

func functionReturn(line string) {

}

func decipherType (args string) (string) {
	//checking string
	condition1 := strings.HasPrefix(args,"\"") || strings.HasPrefix(args,"'")
	condition2 := strings.HasSuffix(args,"\"") || strings.HasSuffix(args,"'")
	if condition1 && condition2 {
		return STRING
	}

	//checking if double
	if strings.Contains(args,".") {
		return DOUBLE
	}

	//checking number
	temp := []rune(args)
	itIsADigit := true
	for i := 0; i < len(temp); i++ {
		itIsADigit = itIsADigit && unicode.IsDigit(temp[i])
	}
	if itIsADigit {
		return NUMBER
	}


	//checking if charge [boolean]
	if args == "true" || args == "false" {
		return BOOLEAN
	}

	return SYSTEM_STRING

}

func createStringRep (line string,TYPE string) (string){
	switch TYPE {
	case STRING:
		return "\u001B[96m"+string([]rune(line)[1:len(line)-1])+"\u001B[0m"
	case SYSTEM_STRING:
		return "\u001B[95m"+line+"\u001B[0m"
	default:
		return "\u001B[96m"+strings.TrimSpace(line)+"\u001B[0m"
	}
}

//compute is incomplete ##########################################3
func (caller ScopeNode) compute(value string) (Data) {
	value =  strings.TrimSpace(value)
	elems := strings.Split(value," ")
	for i := 0; i < len(elems); i++ {
		if (elems[i] != "" || elems[i] != "" ){
			value,found := caller.searchTreeFor(elems[i])
			if found {
				elems[i] = value.value
			}
		}
	}

	value = ""

	for i := 0; i < len(elems)-1; i++ {
		value = value + elems[i] + " "
	}
	value = value + elems[len(elems)-1]
	
	value = getFinalValue(value)

	Type := decipherType(strings.TrimSpace(value))
	valueD := value

	if Type == STRING {
		value = string([]rune(value)[1:len(value)-1])
		valueD = "\""+value+"\""
	}

	return Data{value,Type,createStringRep(valueD,Type)}
}

func getFinalValue (line string) (string) {
	answer := multiSplitOperators(line)
	return answer
}

func multiSplitOperators(line string) (string){
	var splitPoints []int
	mapp := make(map[int]rune)
	elems := []rune(" "+line)

	headPointer := &DataNode{Data{},NodePointer{&DataNode{Data{},NodePointer{},false},' '},true}

	splitPoints = append(splitPoints,0)
	for i := 0; i < len(oprList); i++ {
		for j:=0;j<len(elems);j++ {
			if elems[j] == oprList[i].rep {
				splitPoints = append(splitPoints,j)
				mapp[j] = oprList[i].rep
			}
		}
	}

	splitPoints = append(splitPoints,len(elems))

	var parts []string

	for i := 0; i < len(splitPoints)-1; i++ {
		str := strings.TrimSpace(string(elems[splitPoints[i]+1:splitPoints[i+1]])) + string(mapp[splitPoints[i+1]])
		parts = append(parts,str)
	}

	for i := 0; i < len(parts); i++ {
		symbol := []rune(parts[i])[len(parts[i])-1]
		parts[i] = string([]rune(parts[i])[:len(parts[i])-1])
		headPointer.iterateAndAdd(parts[i],symbol)
	}

	answer := headPointer.calculate()
	return answer
}

func (caller *DataNode) iterateAndAdd (element string,symbol rune) {
	presentNode := &caller
	done := false

	for !done{
		if !((**presentNode).pointer.parent.doesNotExist) {
			dataNode := encapsulate(element,symbol)
			(*presentNode).pointer.parent = &dataNode
			break
		} else {
			presentNode = &(**presentNode).pointer.parent
		}
	}
}

func encapsulate (element string,symbol rune) (DataNode) {
	element = strings.TrimSpace(element)
	Type := decipherType(element)
	data := Data{element,Type,createStringRep(element,Type)}
	pointer := NodePointer{&DataNode{},symbol}
	return DataNode{data,pointer,true}
}

func (caller *DataNode) calculate () (string){
	result := "resnull"
	nextSymbol := ' '
	for ((*(*caller).pointer.parent).pointer.parent.doesNotExist) {
		child := *(*caller).pointer.parent
		parent := (*(*(*caller).pointer.parent).pointer.parent)

		for i := 0; i < len(oprList); i++ {
			if oprList[i].rep == child.pointer.action {
				result,_ = oprList[i].action(child.data,parent.data)
				break
			}
		}

		
		nextSymbol = (*(*caller).pointer.parent).pointer.action
		dataNode := encapsulate(result,nextSymbol)

		dataNode.pointer.parent = parent.pointer.parent

		(*caller).pointer.parent = &dataNode
	}
	return result
}
