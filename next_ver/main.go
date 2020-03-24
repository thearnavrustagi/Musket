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
	NORMAL_ASSIGNMENT string = "="
	LEXICAL_VAR_SCOPE_ASIGNMENT string = ":= "
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
	NOT string = "not "
	IF string = "if "
	IF_ARGS_ARE_EQUAL string = "are equal"
	ELSE string = "else "
	WHILE string = "while "

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
	calledBy *MethodData
	calledAt int
	calledWith string
}

type ScopeNode struct {
	parent *ScopeNode
	varSave map[string]Data
	presentLine int
	owner *MethodData
}

type LoopData struct {
	data MethodData
	condition string
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
var operatorList []Operators
var buldFailure bool

//execution tree
var executionListHeadNode *MethodData
func main() {
	argHandler := InititializeCMD()
	//this function is HUGE and is at the end
	InititializeOperators()
	args := getUserArgs()
	
	if args != "shell" {
		methodSave = make(map[string]*MethodData)
		varSyntaxSave = make(map[string]VarSyntaxData)
		globalScope = ScopeNode{nil,make(map[string]Data),0,&MethodData{}}

		
		program := Interpret(args,argHandler)
		startExec(program)
	} else {
		fmt.Println("WELCOME TO VIPER LANG")
	
		for true {
			methodSave = make(map[string]*MethodData)
			varSyntaxSave = make(map[string]VarSyntaxData)
			globalScope = ScopeNode{nil,make(map[string]Data),0,&MethodData{}}
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

func err(args string) {
	fmt.Println("\u001B[91m"+args+"\u001B[0m")
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
		if varSyntaxIsDeclarable(program[i]) {
			endIndex := declareVarSyntax(program,i)
			program = deleteBlock(program,i,endIndex)
		}
	}
	fmt.Println("\u001B[92mFIRST ASSIGNMENT RUN \nSTATUS: COMPLETE\u001B[0m\n")
	StartExecution(program)
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

	node := MethodData{param,block, ScopeNode{&caller,make(map[string]Data),0,&MethodData{}},&MethodData{},0,""}

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
	if (startIndex == endIndex) {
		return Block{},-1
	}
	snippet := make([]string,endIndex-startIndex)
	for i := startIndex; i <endIndex; i++ {
		snippet[i-startIndex] = program[i]
	}
	return Block{snippet},endIndex
}

func varSyntaxIsDeclarable (args string) (bool){
	return strings.Contains(args,SYNTACTIC_ASSIGNMENT)
}

func declareVarSyntax (program []string,index int) (int){
	name := strings.TrimSpace(strings.Split(program[index],SYNTACTIC_ASSIGNMENT)[0])

	dataBlock,endIndex := getBlock(program,index,'(',')')
	vsData := VarSyntaxData{dataBlock}
	varSyntaxSave[name] = vsData
	varSyntaxNames = append(varSyntaxNames,name)

	return endIndex
}

func StartExecution(program []string) {
	for i := 0; i < len(methodNames); i++ {
		param := methodSave[methodNames[i]].parameters
		data := methodSave[methodNames[i]].data.data
		scopeNode := methodSave[methodNames[i]].scopeNode

		methodSave[methodNames[i]] = &MethodData{param , Block{replaceSpecialSyntax(data)} , scopeNode , &MethodData{},0,""}

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
	data := MethodData{main.parameters,main.data,main.scopeNode,&MethodData{},0,""}
	data.runThrough(0)
	return ""
}

func autoCreateFunctionMAIN (program []string) {
	main := MethodData{"",Block{program},ScopeNode{&globalScope,make(map[string]Data),0,&MethodData{}},&MethodData{},0,""}
	main.data.data = replaceSpecialSyntax(main.data.data)
	methodSave["main"] = &main
	callFunctionMAIN()
}

func (data MethodData) runThroughFnc (index int,scpNode ScopeNode) {
	parts := strings.Split(strings.TrimSpace(data.data.data[0]),"=")
	lhs := strings.Split(parts[0],",")
	rhs := strings.Split(parts[1],",")

	if len(lhs) != len(rhs){
		err("insufficient parameters to call function")
	}

	for i := 0; i < len(lhs); i++ {
		lhs[i],rhs[i] = strings.TrimSpace(lhs[i]),strings.TrimSpace(rhs[i])
		var_data := scpNode.compute(rhs[i])
		data.scopeNode.varSave[lhs[i]] = var_data
	}

	data.runThrough(1)

}

func (data MethodData) runThrough (index int) {
	program := data.data.data
	data.scopeNode.owner = &data

	for i := index; i < len(program); i++ {
		//fmt.Println(program)

		data.scopeNode.presentLine = i
		program[i] = strings.TrimSpace(program[i])

		done,_:= data.scopeNode.checkSpecialFunctions(program[i])
		if done {
			//i = i + increment
			continue
		}
		
		if lexScopingIsBeingRequested(program[i]) {
			data.scopeNode.lexScopingAssign(program[i])
			continue
		}

		if assignable(program[i]) {
			data.scopeNode.assign(program[i])
			continue
		}


		if varSyntaxIsDeclarable(program[i]) {
			endIndex := declareVarSyntax(program,i)
			program = deleteBlock(program,i,endIndex)
		}

		if strings.HasPrefix(program[i],"{") && strings.HasSuffix(program[i],"}") {
			str := string([]rune(program[i])[1:len(program[i])-1])
			val,found := data.scopeNode.searchTreeFor(str)
			value := val.value
			if found {
				program[i] = value
				i = i-1
				continue
			}
		}

		funcCall,methodName := functionCall(program[i])

		if  funcCall {
			data.scopeNode.callFunction(program[i],methodName,data,i)
			continue
		}

		
		if strings.HasPrefix(program[i],RETURN_STATEMENT) {
			data.functionReturn(program[i])
			continue
		}
	}

	data = MethodData{}
}


func assignable(line string) (bool) {
	return strings.Contains(line,NORMAL_ASSIGNMENT)
}

func (caller ScopeNode) assign (args string) {
	if strings.HasPrefix(args,VAR_DECALRATION) {
		args = strings.TrimSpace(string([]rune(args)[len(VAR_DECALRATION):]))
	} else {
		args = strings.TrimSpace(args)
	}

	tempParts := strings.Split(args,NORMAL_ASSIGNMENT)
	lhs := strings.Split(strings.TrimSpace(tempParts[0]),",")
	rhs := strings.Split(strings.TrimSpace(tempParts[1]),",")

	caller.assignValues(lhs,rhs,false)
}

func lexScopingIsBeingRequested(line string) (bool) {
	return strings.Contains(line,LEXICAL_VAR_SCOPE_ASIGNMENT)
}

func (caller ScopeNode) lexScopingAssign (args string) {
	if strings.HasPrefix(args,LEXICAL_VAR_SCOPE_ASIGNMENT) {
		args = strings.TrimSpace(string([]rune(args)[len(LEXICAL_VAR_SCOPE_ASIGNMENT):]))
	} else {
		args = strings.TrimSpace(args)
	}

	tempParts := strings.Split(args,LEXICAL_VAR_SCOPE_ASIGNMENT)
	lhs := strings.Split(strings.TrimSpace(tempParts[0]),",")
	rhs := strings.Split(strings.TrimSpace(tempParts[1]),",")

	caller.assignValues(lhs,rhs,true)
}

func (caller ScopeNode) assignValues (names,values []string,lexScoping bool) {
	for i := 0; i < len(values); i++ {
		names[i] = strings.TrimSpace(names[i])
		val := caller.compute(values[i])
		if lexScoping {
			fmt.Println("lex")
			caller.varSave[names[i]] = val
		} else {
			caller.searchTheTreeAndPlace(names[i],val)
		}
	}
}
func (caller ScopeNode) searchTheTreeAndPlace (name string,value Data) {
	presentNode := &caller
	successFullNode := &caller
	for !reflect.DeepEqual(presentNode.owner,&MethodData{}) {
		
		if (presentNode.varSave[name] != Data{}) {
			successFullNode = presentNode
		}

		presentNode = presentNode.parent
	}

	successFullNode.varSave[name] = value
}
func deleteBlock (program []string,startIndex int,endIndex int) ([]string){
	for i := startIndex; i <= endIndex; i++ {
		program[i] = ""
	}

	return program
}

func functionCall(args string) (bool,string){
	for i := 0; i < len(methodNames); i++ {
		if (strings.HasPrefix(args,methodNames[i])) {
			return true,methodNames[i]
		}
	}
	return false,""
}

func (caller ScopeNode)callFunction (method string,name string,calledBy MethodData,callingAt int) {
	methodSave[name].calledAt = callingAt
	methodSave[name].calledBy = &calledBy
	methodSave[name].calledWith = method

	method = removeSyntacticSugar(strings.TrimSpace(method))
	param := string([]rune(method)[len(name)+1:])

	assignmentString := methodSave[name].parameters+" = "+param


	methodSave[name].data.data[0] = assignmentString
	methodSave[name].runThroughFnc(0,caller)
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

func (caller ScopeNode) checkSpecialFunctions(line string) (bool,int){
	if strings.HasPrefix(line,PRINTING) {
		caller.print(line)
		return true,0
	}

	if strings.HasPrefix(line,PURE_PRINT) {
		caller.pure_print(line)
		return true,0
	}

	if strings.HasPrefix(line,INPUT) {
		caller.takeInput(line)
		return true,0
	}

	if strings.HasPrefix(line,NOT) {
		caller.NOT(line)
		return true,0
	}

	if strings.HasPrefix(line,IF) {
		caller.tryIf(line)
		return true,1
	}

	if strings.HasPrefix(line,WHILE) {
		caller.tryWhile(line)
		return true,1
	}
	return false,0
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
	presentNode := &caller
	
	for !reflect.DeepEqual(presentNode.owner,&MethodData{}) {
		if !reflect.DeepEqual(presentNode.varSave[name],Data{}) {
			return presentNode.varSave[name],true
		}

		presentNode = presentNode.parent
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
	nstr := "\""+string([]rune(str)[:len(str)-1])+"\""

	data := encapsulate(nstr,' ').data

	caller.searchTheTreeAndPlace(parts[0],data)
}

func (caller ScopeNode) NOT (line string) {
	line = strings.TrimSpace(line)
	line = string([]rune(line)[len(NOT):])
	parts := strings.Split(line,",")

	name := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	val,found := caller.searchTreeFor(value)

	if found {
		value = val.value
	}
	fmt.Println("|"+value+"|")

	boolCounterpart,_:=strconv.ParseBool(value)
	boolVal :=  !boolCounterpart
	value = strconv.FormatBool(boolVal)

	var data Data = encapsulate(value,' ').data

	caller.varSave[name] = data
}

func (caller ScopeNode) tryIf (line string) {
	ogLine := line

	line = string([]rune(line)[len(IF):])
	line = strings.TrimSpace(line)
	if !strings.HasSuffix(line,"{") {
		err("if block should always end with \"{\"")
		os.Exit(0)
	}
	line = strings.TrimSpace(string([]rune(line)[:len(line)-1]))
	line = caller.computeBooleanCondition(line)

	data := caller.owner.data.data
	index := caller.presentLine

	//declaring if block
	if_dataBlock,ifEndIndex := getBlock(data,index,'{','}')

	scpNode := ScopeNode{&caller,make(map[string]Data),0,&MethodData{}}
	ifBlock := MethodData{"",if_dataBlock,scpNode,caller.owner,scpNode.presentLine,ogLine}

	//declaring else block
	else_dataBlock,elseEndIndex := getBlock(data,ifEndIndex+1,'{','}')
	scpNode = ScopeNode{&caller,make(map[string]Data),0,&MethodData{}}
	elseBlock := MethodData{"",else_dataBlock,scpNode,caller.owner,scpNode.presentLine+ifEndIndex,ogLine}

	if line == "true" {
		ifBlock.runThrough(1)
		if elseEndIndex == -1 {
			if strings.HasPrefix(ifBlock.calledBy.data.data[0],WHILE) {
				mData := *(caller.owner)

				prevCallIndex := mData.calledBy.scopeNode.presentLine

				line := mData.data.data[0]

				line = string([]rune(line)[len(WHILE):])
				line = strings.TrimSpace(line)

				line = strings.TrimSpace(string([]rune(line)[:len(line)-1]))
				
				condition := line

				caller.owner.calledBy.scopeNode.loop(mData,ifEndIndex,prevCallIndex+len(mData.data.data),condition)
			
			} else {
				caller.owner.runThrough(ifEndIndex)
				os.Exit(0)
			}
		

		} else {
			if strings.HasPrefix(ifBlock.calledBy.data.data[0],WHILE) {
				caller.CallerIsWhile(elseEndIndex)
			} else {
				caller.owner.runThrough(elseEndIndex)
				os.Exit(0)
			}
		}
	} else {
		else_exists,_ := elseExists(data,index)
		if else_exists  {
			elseBlock.runThrough(0)
			if strings.HasPrefix(ifBlock.calledBy.data.data[0],WHILE) {
				caller.CallerIsWhile(elseEndIndex)
			} else {
				caller.owner.runThrough(elseEndIndex)
				os.Exit(0)
			}
		} else {
			if strings.HasPrefix(ifBlock.calledBy.data.data[0],WHILE) {
				caller.CallerIsWhile(ifEndIndex)
			} else {
				caller.owner.runThrough(ifEndIndex)
				os.Exit(0)
			}
		}

	}
}

func (caller ScopeNode)CallerIsWhile (endIndex int) {
		mData := *(caller.owner)

		prevCallIndex := mData.calledBy.scopeNode.presentLine

		line := mData.data.data[0]

		line = string([]rune(line)[len(WHILE):])
		line = strings.TrimSpace(line)

		line = strings.TrimSpace(string([]rune(line)[:len(line)-1]))
				
		condition := line
		caller.owner.calledBy.scopeNode.loop(mData,endIndex,prevCallIndex+len(mData.data.data),condition)
}

func (caller ScopeNode) computeBooleanCondition (line string) (string){
	line = removeParanthesisIfPresent(line)
	
	if strings.HasSuffix(line,IF_ARGS_ARE_EQUAL) {
		line = strings.TrimSpace(string([]rune(line)[:len(line)-len(IF_ARGS_ARE_EQUAL)]))
		line = removeParanthesisIfPresent(line)
		line = removeSyntacticSugar(line)
		parts := strings.Split(line,",")


		for i := 0; i < len(parts); i++ {
			parts[i] = strings.TrimSpace(parts[i])
			temp := caller.compute(parts[i])
			parts[i] = temp.value
		}
		equal := true
		for i := 0; i < len(parts)-1; i++ {
			equal = equal && (parts[i] == parts[i+1])			
		}
		line = strconv.FormatBool(equal)
	} else {
		line = caller.compute(line).value
	}

	return line
}

func elseExists (program []string,index int) (bool,int){
	for i := index; i<len(program); i++ {
		if strings.HasPrefix(strings.TrimSpace(program[i]),ELSE) {
			return true,i
		}
	}
	return false,0
}

func  deleteElse (program []string,index int) {
	_,endIndex := getBlock(program,index,'{','}')
	deleteBlock(program,index,endIndex)
}

func removeParanthesisIfPresent (line string) (string) {
	if strings.HasPrefix(line,"(")&&strings.HasSuffix(line,")") {
		line = string([]rune(line)[1:len(line)-1])
	}

	return line
}

func (caller ScopeNode) tryWhile (line string) {
	ogLine := line
	index := caller.presentLine
	program := caller.owner.data.data

	line = string([]rune(line)[len(WHILE):])
	line = strings.TrimSpace(line)

	if !strings.HasSuffix(line,"{") {
		err("while block should always end with \"{\"")
		os.Exit(0)
	}
	line = strings.TrimSpace(string([]rune(line)[:len(line)-1]))
	
	condition := line
	//while block creation
	while_dataBlock,endIndex := getBlock(program,index,'{','}')

	scpNode := ScopeNode{&caller,make(map[string]Data),0,&MethodData{}}
	WhileMethod := MethodData{"",while_dataBlock,scpNode,caller.owner,scpNode.presentLine,ogLine}


	caller.loop(WhileMethod,1,endIndex,condition)
}

func (caller ScopeNode) loop (loopMethod MethodData,startIndex,endIndex int,condition string) {
	value := caller.computeBooleanCondition(condition)

	//fmt.Println(condition,endIndex)
	if value == "true" {
		loopMethod.runThrough(startIndex)
		caller.loop(loopMethod,1,endIndex,condition)
	} else {
		if strings.HasPrefix(loopMethod.calledBy.data.data[0],WHILE) {
			mData := *(caller.owner)

			prevCallIndex := mData.calledBy.scopeNode.presentLine

			line := mData.data.data[0]

			line = string([]rune(line)[len(WHILE):])
			line = strings.TrimSpace(line)

			line = strings.TrimSpace(string([]rune(line)[:len(line)-1]))
			
			condition := line

			caller.owner.calledBy.scopeNode.loop(mData,endIndex,prevCallIndex+len(mData.data.data),condition)
			os.Exit(0)
		} else {
			loopMethod.calledBy.runThrough(endIndex)
			os.Exit(0)
		}
	}
}

func (data MethodData) functionReturn(line string) {
	caller := data.scopeNode.owner.scopeNode

	returnable := string([]rune(line)[len(RETURN_STATEMENT):])
	returnable = strings.TrimSpace(returnable)
	dataNode,found := caller.searchTreeFor(returnable)
	var returnData Data
	
	returnData = encapsulate(returnable,' ').data

	if found {
		returnData = dataNode 
	}

	data.replaceFunctionCallForValue(returnData)
}

func (data MethodData) replaceFunctionCallForValue (returnData Data) {
	caller := data.calledBy
	lines := caller.data.data
	index := data.calledAt
	line := lines[index]
	value := returnData.value
	line = strings.Replace(line,data.calledWith,value,-1)
	data.calledBy.data.data[index] = line
	caller.runThrough(index)
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
	if args == "true"||args == "plus"||args == "minus"||args == "false" {
		return BOOLEAN
	}

	return SYSTEM_STRING

}

func createStringRep (line string,TYPE string) (string){
	switch TYPE {
	case STRING:
		return "\u001B[96m"+string([]rune(line)[1:len(line)-1])
	case BOOLEAN:
		if line == "plus" || line == "true"{
			return "true"
		} else {
			return "false"
		}
	case SYSTEM_STRING:
		return "\u001B[95m"+line
	default:
		return "\u001B[96m"+strings.TrimSpace(line)
	}
}

func (caller ScopeNode) compute(value string) (Data) {
	value = strings.TrimSpace(value)
	if hasNoOperators(value) {
		value = value+"+0"
		/*temp,found := caller.searchTreeFor(value)
		if found {
			value = temp.value
		}
		
		Type := decipherType(strings.TrimSpace(value))
		valueD := value
		
		return Data{value,Type,createStringRep(valueD,Type)}*/
	}

	value = caller.applyOperators(value)

	Type := decipherType(strings.TrimSpace(value))
	valueD := value
	
	return Data{value,Type,createStringRep(valueD,Type)}
}

func hasNoOperators(args string) (bool){
	returnable := true

	for i := 0; i < len(operatorList); i++ {
		returnable = returnable && !strings.Contains(args,string(operatorList[i].rep))
	}

	return returnable
}

func (caller ScopeNode) applyOperators(line string) (string){
	var splitPoints []int
	mapp := make(map[int]rune)
	elems := []rune(" "+line)

	headPointer := &DataNode{Data{},NodePointer{&DataNode{Data{},NodePointer{},false},' '},true}

	splitPoints = append(splitPoints,0)
	for i := 0; i < len(operatorList); i++ {
		for j:=0;j<len(elems);j++ {
			if elems[j] == operatorList[i].rep {

				splitPoints = append(splitPoints,j)
				mapp[j] = operatorList[i].rep
			}
		}
	}

	splitPoints = append(splitPoints,len(elems))

	var parts []string

	for i := 0; i < len(splitPoints)-1; i++ {
		repr :=strings.TrimSpace(string(elems[splitPoints[i]+1:splitPoints[i+1]]))
		val,found := caller.searchTreeFor(repr)

		if found {
			str := val.value + string(mapp[splitPoints[i+1]])
			parts = append(parts,str)
		} else {
			isAFunction,name := checkIfAFunction(repr)
			if isAFunction {
				caller.callFunction(repr,name,*caller.owner,caller.presentLine)
			}
			str := repr + string(mapp[splitPoints[i+1]])
			parts = append(parts,str)
		}
	}

	for i := 0; i < len(parts); i++ {
		symbol := []rune(parts[i])[len(parts[i])-1]
		parts[i] = string([]rune(parts[i])[:len(parts[i])-1])
		headPointer.iterateAndAdd(parts[i],symbol)
	}

	answer := headPointer.calculate()
	return answer
}

func checkIfAFunction (args string) (bool,string){
	return functionCall(args)
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
	rep := createStringRep(element,Type)
	if Type == STRING{
		element = string([]rune(element)[1:len(element)-1])
	}
	data := Data{element,Type,rep}
	pointer := NodePointer{&DataNode{},symbol}
	return DataNode{data,pointer,true}
}

func (caller *DataNode) calculate () (string){
	result := "\u001B[91m"+NULL+"\u001B[0m"
	nextSymbol := ' '
	for ((*(*caller).pointer.parent).pointer.parent.doesNotExist) {
		child := *(*caller).pointer.parent
		parent := (*(*(*caller).pointer.parent).pointer.parent)

		for i := 0; i < len(operatorList); i++ {
			if operatorList[i].rep == child.pointer.action {
				result,_ = operatorList[i].action(child.data,parent.data)
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

//operators
func InititializeOperators() {
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
			
				return ("\""+arg1.value+arg2.value+"\""),true			
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
				return ("\""+arg1.value+arg2.value+"\""),true
			
			}
		
		} else if arg1.Type == BOOLEAN {

			if arg2.Type == NUMBER {
		
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
		
			} else if arg2.Type == DOUBLE {
		
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
		
			} else if arg2.Type == BOOLEAN {
				var1,var2 := arg1.value,arg2.value
				if var1 == "plus" {
					var1 = "true"
				} else if var1 == "minus" {
					var1 = "false"
				}

				if var2 == "plus" {
					var2 = "true"
				} else if var2 == "minus" {
					var2 = "false"
				}
		
				val1,_ := strconv.ParseBool(var1)
				val2,_ := strconv.ParseBool(var2)
				return strconv.FormatBool(val1&&val2),true
		
			} else {
				return ("\""+arg1.value+arg2.value+"\""),true		
			}
		
		} else {
			return ("\""+arg1.value+arg2.value+"\""),true
		
		}
		
		return NULL,true
	}}

	MINUS := Operators{'-',
	func(arg1 Data,arg2 Data)(string,bool){
		if arg1.Type == NUMBER {

			if arg2.Type == NUMBER {
				
				num1,_ := strconv.Atoi(arg1.value)
				num2,_ := strconv.Atoi(arg2.value)
				return strconv.Itoa((num1-num2)),true
			
			} else if arg2.Type == DOUBLE {
				
				dec := arg1.value+".0"
				dec1,_ := strconv.ParseFloat(dec,64)
				dec2,_ := strconv.ParseFloat(arg2.value,64)
				return strconv.FormatFloat(dec1-dec2,'f',-1,64),true
			
			} else if arg2.Type == BOOLEAN {
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
			
			} else {
				r1,r2 := []rune(arg1.value),[]rune(arg2.value)
				sum1,sum2 :=0,0
				for i := 0; i < len(r1); i++ {
					sum1 = sum1 + int(r1[i])
				}
				for i := 0; i < len(r2); i++ {
					sum2 = sum2 + int(r2[i])
				}
				if sum2>sum1 {
					ans := arg2.value 
					return "\""+ans+"\"",true
				} else {
					ans := arg1.value
					return "\""+ans+"\"",true
				}
			
				return arg1.value+arg2.value,true
			}
		
		} else if arg1.Type == DOUBLE {
			
			if arg2.Type == DOUBLE {
			
				num1,_ := strconv.ParseFloat(arg1.value,64)
				num2,_ := strconv.ParseFloat(arg2.value,64)
				return strconv.FormatFloat(num1-num2,'f',-1,64),true
			
			} else if arg2.Type == NUMBER {
			
				dec := arg2.value+".0"
				dec1,_ := strconv.ParseFloat(arg1.value,64)
				dec2,_ := strconv.ParseFloat(dec,64)
				return strconv.FormatFloat(dec1-dec2,'f',-1,64),true
			
			} else if arg2.Type == BOOLEAN {
			
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
			
			} else {
				r1,r2 := []rune(arg1.value),[]rune(arg2.value)
				sum1,sum2 :=0,0
				for i := 0; i < len(r1); i++ {
					sum1 = sum1 + int(r1[i])
				}
				for i := 0; i < len(r2); i++ {
					sum2 = sum2 + int(r2[i])
				}
				if sum2>sum1 {
					ans := arg2.value 
					return "\""+ans+"\"",true
				} else {
					ans := arg1.value
					return "\""+ans+"\"",true
				}
			
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
				var1,var2 := arg1.value,arg2.value
				if var1 == "plus" {
					var1 = "true"
				} else if var1 == "minus" {
					var1 = "false"
				}

				if var2 == "plus" {
					var2 = "true"
				} else if var2 == "minus" {
					var2 = "false"
				}
		
				val1,_ := strconv.ParseBool(var1)
				val2,_ := strconv.ParseBool(var2)
				return strconv.FormatBool(val1&&(!val2)),true
		
			} else {
				r1,r2 := []rune(arg1.value),[]rune(arg2.value)
				sum1,sum2 :=0,0
				for i := 0; i < len(r1); i++ {
					sum1 = sum1 + int(r1[i])
				}
				for i := 0; i < len(r2); i++ {
					sum2 = sum2 + int(r2[i])
				}
				if sum2>sum1 {
					ans := arg2.value 
					return "\""+ans+"\"",true
				} else {
					ans := arg1.value
					return "\""+ans+"\"",true
				}
			
				return arg1.value+arg2.value,true
		
			}
		
		} else {
			r1,r2 := []rune(arg1.value),[]rune(arg2.value)
			sum1,sum2 :=0,0
			for i := 0; i < len(r1); i++ {
				sum1 = sum1 + int(r1[i])
			}
			for i := 0; i < len(r2); i++ {
				sum2 = sum2 + int(r2[i])
			}
			if sum2>sum1 {
				ans := arg2.value 
				return "\""+ans+"\"",true
			} else {
				ans := arg1.value
				return "\""+ans+"\"",true
			}
			
			return NULL,true
		}
		
		return NULL,true
	}}

	MULT := Operators{'*',
	func(arg1 Data,arg2 Data)(string,bool){
		if arg1.Type == NUMBER {

			if arg2.Type == NUMBER {
				
				num1,_ := strconv.Atoi(arg1.value)
				num2,_ := strconv.Atoi(arg2.value)
				return strconv.Itoa((num1*num2)),true
			
			} else if arg2.Type == DOUBLE {
				
				dec := arg1.value+".0"
				dec1,_ := strconv.ParseFloat(dec,64)
				dec2,_ := strconv.ParseFloat(arg2.value,64)
				return strconv.FormatFloat(dec1*dec2,'f',-1,64),true
			
			} else if arg2.Type == BOOLEAN {
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
			
			} else {
				times := len(arg2.value)
				ret := ""
				for i := 0; i < times; i++ {
					ret = ret+arg1.value
				}
				return ("\""+ret+"\""),true
			}
		
		} else if arg1.Type == DOUBLE {
			
			if arg2.Type == DOUBLE {
			
				num1,_ := strconv.ParseFloat(arg1.value,64)
				num2,_ := strconv.ParseFloat(arg2.value,64)
				return strconv.FormatFloat(num1*num2,'f',-1,64),true
			
			} else if arg2.Type == NUMBER {
			
				dec := arg2.value+".0"
				dec1,_ := strconv.ParseFloat(arg1.value,64)
				dec2,_ := strconv.ParseFloat(dec,64)
				return strconv.FormatFloat(dec1*dec2,'f',-1,64),true
			
			} else if arg2.Type == BOOLEAN {
			
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
			
			} else {
				times := len(arg2.value)
				ret := ""
				for i := 0; i < times; i++ {
					ret = ret+arg1.value
				}
				return ("\""+ret+"\""),true
			}
		
		} else if arg1.Type == BOOLEAN {

			if arg2.Type == NUMBER {
		
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
		
			} else if arg2.Type == DOUBLE {
		
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
		
			} else if arg2.Type == BOOLEAN {
				var1,var2 := arg1.value,arg2.value
				if var1 == "plus" {
					var1 = "true"
				} else if var1 == "minus" {
					var1 = "false"
				}

				if var2 == "plus" {
					var2 = "true"
				} else if var2 == "minus" {
					var2 = "false"
				}
		
				val1,_ := strconv.ParseBool(var1)
				val2,_ := strconv.ParseBool(var2)
				return strconv.FormatBool(val1&&val2),true
		
			} else {times := len(arg2.value)
				ret := ""
				for i := 0; i < times; i++ {
					ret = ret+arg1.value
				}
				return ("\""+ret+"\""),true
		
			}
		
		} else {
			if arg2.Type == NUMBER {
				ret := ""
				times,_ := strconv.Atoi(arg2.value)

				for i := 0; i < times; i++ {
					ret = ret+arg1.value
				}

				return ("\""+ret+"\""),true
			
			}else{
				times := len(arg2.value)
				ret := ""
				for i := 0; i < times; i++ {
					ret = ret+arg1.value
				}
				return ("\""+ret+"\""),true
			}
		}
		
		return NULL,true
	}}

	DIVIDE := Operators{'/',
		func(arg1 Data,arg2 Data)(string,bool){
			if arg1.Type == NUMBER {

				if arg2.Type == NUMBER {
					
					num1,_ := strconv.Atoi(arg1.value)
					num2,_ := strconv.Atoi(arg2.value)
					return strconv.Itoa((num1/num2)),true
				
				} else if arg2.Type == DOUBLE {
					
					dec := arg1.value+".0"
					dec1,_ := strconv.ParseFloat(dec,64)
					dec2,_ := strconv.ParseFloat(arg2.value,64)
					return strconv.FormatFloat(dec1/dec2,'f',-1,64),true
				
				} else if arg2.Type == BOOLEAN {
					fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
					os.Exit(0)
				
				} else {
					num1 := len(arg1.value)
					num2 := len(arg2.value)
					result := num1 - num2
					res := ""

					if (result < 0) {
						result = num2 - num1
						rnu := []rune(arg2.value)
						res = string(rnu[:len(arg2.value)-len(arg1.value)]) 
					} else {
						rnu := []rune(arg1.value)
						res = string(rnu[:len(arg1.value)-len(arg2.value)])
					}

					
					return ("\""+res+"\""),true
				}
			
			} else if arg1.Type == DOUBLE {
				
				if arg2.Type == DOUBLE {
				
					num1,_ := strconv.ParseFloat(arg1.value,64)
					num2,_ := strconv.ParseFloat(arg2.value,64)
					return strconv.FormatFloat(num1/num2,'f',-1,64),true
				
				} else if arg2.Type == NUMBER {
				
					dec := arg2.value+".0"
					dec1,_ := strconv.ParseFloat(arg1.value,64)
					dec2,_ := strconv.ParseFloat(dec,64)
					return strconv.FormatFloat(dec1/dec2,'f',-1,64),true
				
				} else if arg2.Type == BOOLEAN {
				
					fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
					os.Exit(0)
				
				} else {
					num1 := len(arg1.value)
					num2 := len(arg2.value)
					result := num1 - num2
					res := ""

					if (result < 0) {
						result = num2 - num1
						rnu := []rune(arg2.value)
						res = string(rnu[:len(arg2.value) -len(arg1.value)]) 
					} else {
						rnu := []rune(arg1.value)
						res = string(rnu[:len(arg1.value)-len(arg2.value)])
					}

					
					return ("\""+res+"\""),true
				
				}
			
			} else if arg1.Type == BOOLEAN {

				if arg2.Type == NUMBER {
			
					fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
					os.Exit(0)
			
				} else if arg2.Type == DOUBLE {
			
					fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
					os.Exit(0)
			
				} else if arg2.Type == BOOLEAN {
					var1,var2 := arg1.value,arg2.value
					if var1 == "plus" {
						var1 = "true"
					} else if var1 == "minus" {
						var1 = "false"
					}

					if var2 == "plus" {
						var2 = "true"
					} else if var2 == "minus" {
						var2 = "false"
					}
			
					val1,_ := strconv.ParseBool(var1)
					val2,_ := strconv.ParseBool(var2)
					return strconv.FormatBool(val1&&(!val2)),true
			
				} else {

					num1 := len(arg1.value)
					num2 := len(arg2.value)
					result := num1 - num2
					res := ""

					if (result < 0) {
						result = num2 - num1
						rnu := []rune(arg2.value)
						res = string(rnu[:len(arg2.value)-len(arg1.value)]) 
					} else {
						rnu := []rune(arg1.value)
						res = string(rnu[:len(arg1.value)-len(arg2.value)])
					}

					
					return ("\""+res+"\""),true
				}
			
			} else {
				if arg2.Type == NUMBER{
					num,_ := strconv.Atoi(arg2.value)
					str := arg1.value
					res := ""
					if num > len(str) {
						err("Too big a number to use / operator on [operands are string and number]")
						os.Exit(0)
					} else {
						rnu := []rune(str)
						res = string(rnu[:len(rnu)-num])
					}
					return ("\""+res+"\""),true
				}else{
					num1 := len(arg1.value)
					num2 := len(arg2.value)
					result := num1 - num2
					res := ""
					if (result < 0) {
						result = num2 - num1
						rnu := []rune(arg2.value)
						res = string(rnu[:len(arg2.value) - len(arg1.value)]) 
					} else {
						rnu := []rune(arg1.value)
						res = string(rnu[:len(arg1.value)-len(arg2.value)])
					}
					return ("\""+res+"\""),true
			}
			
		}
			return NULL,true

	}}
	MOD := Operators{'%',
	func(arg1 Data,arg2 Data)(string,bool){
		if arg1.Type == NUMBER {

			if arg2.Type == NUMBER {
				
				num1,_ := strconv.Atoi(arg1.value)
				num2,_ := strconv.Atoi(arg2.value)
				return strconv.Itoa((num1%num2)),true
			
			} else if arg2.Type == DOUBLE {
				err("DOUBLES do not posses the % opperator")
			
			} else if arg2.Type == BOOLEAN {
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
			
			} else {
				num1 := len(arg1.value)
				num2 := len(arg2.value)
				result := num1 - num2
				res := ""
				if (result < 0) {
					result = num2 - num1
					rnu := []rune(arg2.value)
					res = string(rnu[len(arg2.value) - len(arg1.value):]) 
				} else {
					rnu := []rune(arg1.value)
					res = string(rnu[len(arg1.value)-len(arg2.value):])
				}
				return ("\""+res+"\""),true
			}
		
		} else if arg1.Type == DOUBLE {
			
			if arg2.Type == DOUBLE {

				err("DOUBLES do not posses the % opperator")
			
			} else if arg2.Type == NUMBER {

				err("DOUBLES do not posses the % opperator")
			
			} else if arg2.Type == BOOLEAN {
			
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
			
			} else {
				num1 := len(arg1.value)
				num2 := len(arg2.value)
				result := num1 - num2
				res := ""
				if (result < 0) {
					result = num2 - num1
					rnu := []rune(arg2.value)
					res = string(rnu[len(arg2.value) - len(arg1.value):]) 
				} else {
					rnu := []rune(arg1.value)
					res = string(rnu[len(arg1.value)-len(arg2.value):])
				}
				return ("\""+res+"\""),true
			}
		
		} else if arg1.Type == BOOLEAN {

			if arg2.Type == NUMBER {
		
				fmt.Println("\u001B[96m illegal conversion of charge to number\u001B[0m")
				os.Exit(0)
		
			} else if arg2.Type == DOUBLE {
		
				fmt.Println("\u001B[96m illegal conversion of charge to Double\u001B[0m")
				os.Exit(0)
		
			} else if arg2.Type == BOOLEAN {
				err("Charges do not possess [%] operator")
			} else {
				num1 := len(arg1.value)
				num2 := len(arg2.value)
				result := num1 - num2
				res := ""
				if (result < 0) {
					result = num2 - num1
					rnu := []rune(arg2.value)
					res = string(rnu[len(arg2.value) - len(arg1.value):]) 
				} else {
					rnu := []rune(arg1.value)
					res = string(rnu[len(arg1.value)-len(arg2.value):])
				}
				return ("\""+res+"\""),true	
			}
		
		} else {
			if arg2.Type == NUMBER{
					num,_ := strconv.Atoi(arg2.value)
					str := arg1.value
					res := ""
					if num > len(str) {
						err("Too big a number to use % operator on [operands are string and number]")
						os.Exit(0)
					} else {
						rnu := []rune(str)
						res = string(rnu[len(rnu)-num:])
					}
					return ("\""+res+"\""),true
				}else{
					num1 := len(arg1.value)
					num2 := len(arg2.value)
					result := num1 - num2
					res := ""
					if (result < 0) {
						result = num2 - num1
						rnu := []rune(arg2.value)
						res = string(rnu[len(arg2.value) - len(arg1.value):]) 
					} else {
						rnu := []rune(arg1.value)
						res = string(rnu[len(arg1.value)-len(arg2.value):])
					}
					return ("\""+res+"\""),true
				}
			}
		
		return NULL,true
	}}

	AND := Operators{'&',
	func (arg1,arg2 Data) (string,bool) {
		if arg1.Type == NUMBER {
			if arg2.Type == NUMBER {
				num1,_ := strconv.Atoi(arg1.value)
				num2,_ := strconv.Atoi(arg2.value)
				return strconv.Itoa((num1&num2)),true
			} else {
				err("Both operands should be either of type number or boolean for bitwise/boolean \"and\" to work")
			}
		} else {
			err("Both operands should be either of type number or boolean for bitwise/boolean \"and\" to work")
		}
		if arg1.Type == BOOLEAN {
			if arg2.Type == BOOLEAN {
				var1,var2 := arg1.value,arg2.value
				if var1 == "plus" {
					var1 = "true"
				} else if var1 == "minus" {
					var1 = "false"
				}

				if var2 == "plus" {
					var2 = "true"
				} else if var2 == "minus" {
					var2 = "false"
				}
		
				val1,_ := strconv.ParseBool(var1)
				val2,_ := strconv.ParseBool(var2)
				return strconv.FormatBool(val1&&val2),true
			} else {
				err("the boolean/bitwise \"and\" operator can only be used with boolean/Number operands")
			}
		} else {
			err("the boolean/bitwise \"and\" operator can only be used with boolean/Number operands")
		}
		return NULL,true
	}}

	OR := Operators {'|',
	func (arg1,arg2 Data) (string,bool) {

		if arg1.Type == NUMBER {
			if arg2.Type == NUMBER {
				num1,_ := strconv.Atoi(arg1.value)
				num2,_ := strconv.Atoi(arg2.value)
				return strconv.Itoa((num1|num2)),true
			} else {
				err("Both operands should be either of type number or boolean for bitwise/boolean \"or\" to work")
			}
		} else if arg1.Type == BOOLEAN {
			if arg2.Type == BOOLEAN {
				var1,var2 := arg1.value,arg2.value
				if var1 == "plus" {
					var1 = "true"
				} else if var1 == "minus" {
					var1 = "false"
				}

				if var2 == "plus" {
					var2 = "true"
				} else if var2 == "minus" {
					var2 = "false"
				}
		
				val1,_ := strconv.ParseBool(var1)
				val2,_ := strconv.ParseBool(var2)
				return strconv.FormatBool(val1||val2),true
			} else {
				err("the boolean \"or\" operator can only be used with boolean operands")
			}
		} else if arg1.Type == STRING{ //or gives the greater of the two values
			if arg2.Type == STRING {
				r1,r2 := []rune(arg1.value),[]rune(arg2.value)
				sum1,sum2 :=0,0
				for i := 0; i < len(r1); i++ {
					sum1 = sum1 + int(r1[i])
				}
				for i := 0; i < len(r2); i++ {
					sum2 = sum2 + int(r2[i])
				}
				if sum2>sum1 {
					ans := arg2.value 
					return "\""+ans+"\"",true
				} else {
					ans := arg1.value
					return "\""+ans+"\"",true
				}
			
				return arg1.value+arg2.value,true
			} else {
			err("the given operator is not available in "+arg2.Type+" variables")
			}	
		} else {
			err("the boolean \"or\" operator can only be used with boolean operands")
		}
		return NULL,true
	}}

	XOR := Operators {'^',
	func (arg1,arg2 Data) (string,bool) {

		if arg1.Type == NUMBER {
			if arg2.Type == NUMBER {
				num1,_ := strconv.Atoi(arg1.value)
				num2,_ := strconv.Atoi(arg2.value)
				return strconv.Itoa((num1^num2)),true
			} else {
				err("Both operands should be either of type number or boolean for bitwise/boolean \"xor\" to work")
			}
		} else if arg1.Type == BOOLEAN {
			if arg2.Type == BOOLEAN {
				var1,var2 := arg1.value,arg2.value
				if var1 == "plus" {
					var1 = "true"
				} else if var1 == "minus" {
					var1 = "false"
				}

				if var2 == "plus" {
					var2 = "true"
				} else if var2 == "minus" {
					var2 = "false"
				}
		
				val1,_ := strconv.ParseBool(var1)
				val2,_ := strconv.ParseBool(var2)
				return strconv.FormatBool((val1 || val2) && !(val1 && val2)),true
			} else {
				err("the boolean \"xor\" operator can only be used with boolean operands")
			}
		} else if arg1.Type == STRING { //xor gives the lower of the two values
			if arg2.Type == STRING {
				r1,r2 := []rune(arg1.value),[]rune(arg2.value)
				sum1,sum2 :=0,0
				for i := 0; i < len(r1); i++ {
					sum1 = sum1 + int(r1[i])
				}
				for i := 0; i < len(r2); i++ {
					sum2 = sum2 + int(r2[i])
				}
				if sum2<sum1 {
					ans := arg2.value 
					return "\""+ans+"\"",true
				} else {
					ans := arg1.value
					return "\""+ans+"\"",true
				}
			
				return arg1.value+arg2.value,true
			}
		} else {
			err("the given operator is not available in "+arg1.Type+" variables")
		}
		return NULL,true
	}}

	GREATER := Operators{'>',
	func (arg1,arg2 Data)(string,bool){
		if arg1.Type == NUMBER && arg2.Type == NUMBER {
			num1,_ := strconv.Atoi(arg1.value)
			num2,_ := strconv.Atoi(arg2.value)
			return strconv.FormatBool(num1>num2),true
		} else if arg1.Type == DOUBLE && arg2.Type == DOUBLE {
			num1,_ := strconv.ParseFloat(arg1.value,64)
			num2,_ := strconv.ParseFloat(arg2.value,64)
			return strconv.FormatBool(num1>num2),true
		} else {
			err("invalid operands")
			os.Exit(0)
		}
		return strconv.FormatBool(arg1.value>arg2.value),true
	}}

	LESSER := Operators{'<',
	func (arg1,arg2 Data)(string,bool) {
		if arg1.Type == NUMBER && arg2.Type == NUMBER {
			num1,_ := strconv.Atoi(arg1.value)
			num2,_ := strconv.Atoi(arg2.value)
			return strconv.FormatBool(num1<num2),true
		} else if arg1.Type == DOUBLE && arg2.Type == DOUBLE {
			num1,_ := strconv.ParseFloat(arg1.value,64)
			num2,_ := strconv.ParseFloat(arg2.value,64)
			return strconv.FormatBool(num1<num2),true
		} else {
			err("invalid operands")
			os.Exit(0)
		}
		return strconv.FormatBool(arg1.value<arg2.value),true	
	}}

	INEQUAL := Operators{'!',
	func (arg1,arg2 Data) (string,bool) {
		return strconv.FormatBool(arg1.value!=arg2.value),true
	}}

	EQUAL := Operators{'?',
	func (arg1,arg2 Data) (string,bool) {
		return strconv.FormatBool(arg1.value==arg2.value),true
	}}

	operatorList = append(operatorList,PLUS,MINUS,MULT,DIVIDE,MOD)
	operatorList = append(operatorList,AND,OR,XOR)
	operatorList = append(operatorList,GREATER,LESSER,INEQUAL,EQUAL)
}