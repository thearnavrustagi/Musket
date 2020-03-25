# VIPER
## VIPER SYNTAX (or general syntax)<br>
code is in developement stage please wait for a release
 The language is as follows
```python
__________VIPER_____________
x = 20
a,b,c = 10,15,17
x = 30
test = "x - y"
something

print "okay a is {a}|{b}|{c}|{x}"
	
var x = subtract a,b
var y = add a and b

var something <- (
	print "hello"
	var = {yulo}
)

func subtract (a,b){
	return a-b
}

func add (a,b){
	return a+b
}
 ```
 java like counterpart
 ```java
 ___________JAVA____________
 public static void main() {
 	int x=20;
  	int a = 10,b=15,c=17;
  	x = 30;
	
	//no equivalent syntax for something
	
	System.out.println("okay a is "+a+"|"+b+"|"+c+"|"+x)
	
	int temp = add(a,b)
 }
 
 public void add(int a,int b) {
 	return a+b
 }
 ```
 #### the whole syntax is not made but here are the parts that are made<br>
 PRESENT SYNTAX <br>
 ## RELEASE 1
 - [x] SYNTAX<br>
   - [x] normal assignment
   - [x] syntax assignment
   - [X] printing                      
   - [x] user input                
   - [x] if/else                       
   - [x] while
- [x] OPERATIONS<br>
  - [x] arithmatic                    
  - [x] boolean/logical                       
  - [x] relational                    
  - [x] bitwise                       
- [x] PROGRAMMING FEATURES<br>
  - [x] Lexical Variable Scoping
  - [x] Procedural programming   
  - [x] Functional Programming   
  - [x] Variable Syntax
  - [x] Auto Main
- [x] ERROR_HANDELING
## RELEASE 2
- [ ] SYNTAX
  - [x] Multi Threading (more like processes)
  - [ ] Structures/classes
  - [ ] For loops
  - [ ] flow control and other advanced functions
  - [x] Allowing the usage of parenthesis and semi colons in code
  - [x] Debugging the if/else and while's multiprinting error
- [ ] Maths
  - [ ] priortizing parenthesis
  - [ ] BODMAS
- [ ] Programming features
  - [ ] Smart Interpreter
## Unique Features and Syntax
Well the one which I just developed is called var syntax
- ### VAR_SYNTAX<br>
  well this is basically that you assign a variable like this
  ```python
  var [var_name] <- (
  	# body
  )
  ```
  and this should be outside a method<br>
  to call this you just have to write its name in the main method like so
  ```python
  [var_name]
  ```
  on being called the place where the name was will be changed by its body now why I say changes well let me give you an example
  ```python
  var a,b = 20,30
  add
  print "the sum of a and b is {c}"
  
  var add <- (
  	var c = a + b 
  )
  ```
  yes this is eligible syntax
- ### Automatic main Creation
  Well now there is no need to define a main method all you have to do is write code like you do in pyhon the interpreter will automatically create a main function for you if there is not one now this is where my language overtakes python let me explain in detail<br><br>
  When I was doing python there were two things that did not sync with me one was the fact that no blocks existed like no curly braces just tabs and that did no appeal to me next was the fact that you had to put the "def" block before actually calling the function so due to this the code was untidy and was harder to debug<br>now to fix it I implemented a different way of interpretting the file and that is the three run way<br>This causes speed depreciation but better execution now let me explain every run<br>
  - #### RUN 1
    In this run the methods,global variables and variable syntax is defined and stored
  - #### RUN 2
    In this run the variable syntax is exchanged in every method 
  - #### RUN 3
    This is the final execution in this method main is called and if it does not exist then the remenent of the program with the methods and variable syntax excluded is converted into the main method <br>
  
  Thus it is possible for you to write python's informal syntax in this language but also c++'s formal syntax both are welcomed with open hands
  example
  ```python
  print "hello world"
  printf "hello" and "world"
  func printf (name1,name2){
  	pure_print("{name1} {name2}")
  }
  ```
  or
  ```python
  func main {
  	print "hello world"
  	printf "hello" and "world"	
  }
  func printf = (name1,name2){
  	pure_print("{name1} {name2}")
  }
  ```
  both are eligible syntax
- ### Function Calling and declaration
  **Declaration** is very simple and like golang
  ```python
  func [name] ([args]) {
  	[body]
  }
  ```
  now an important thing to note is the fact that the starting curly brace or '{' should always be **AFTER** the ([args]) or else the code will not execute another thing is that if you want to give no arguments to some function like main you can ommit the ([args]) part just like this 
  ```python
  func [name] {
   [body]
  }
  ```
  The interpreter will understand that you want to give no arguments<br>
  **Function Calling** is another thing that I feel is beautiful about the language why so? well here you go 
  ```python
  func main{
  	var a,b = 20,30
  	println "hello,world",num
  	var x = add a and b
  	var c = choose a or b
  }
  func println(str,num){
  	print "{str}{num}"
  }
  func add (arg1,arg2){
  	ans = arg1+arg2
  	return ans
  }
  func choose (arg1,arg2)(
  	if arg1>arg2 {
		return arg1
	}
	return arg2
  )
  ```
  the first declaration with the comma is the official declaration the rest are syntactic sugar to beautify the code you can add or remove them in the methodSugar.txt file with the "|" delimeter<br>Another thing to remember is that when you call functions with no parameters you have to use "()" like so
  ```python
  //main
  display()
  ```
  assuming that their is some function that requires no parameters
- ### Multi-threading
  Viper supports multi threading the syntax is as follows
  ```python
  run [ [function name] ] concurrently
  ```
  the square brackets can be ommited so this is also eligible syntax
  ```python
  run [function name] concurrently
  ```
  so thats it
- ### Scoping
  lexical variable scope assignment syntax that is 
  ```go
  [varname , [varname] ... ]  := [values,[values] ... ]
  ```
  the ":=" signifies that the assignment will be temporary and when the main thread leaves the block then the values will become default
  ```python
  func main {
  	str = "hello"
	bool = true
	
	if bool {
		str := "bye"
		print "{str}{str}"
	}
	print "{str},world"
  }
  ```
  in this code the output will be 
  ```shell
  byebye
  hello,world
  ```
  so you can create temporary assignments in the code
- ### Variable Declarations
  I have added a formal variable declaration syntax in the language that is 
  ```python
  var [name,[name]....] = [value,[value]....]
  ```
  but the preffered way should be 
  ```python
  [name,[name]....] = [value,[value]....]
  ```
  now the interpreter accepts two ways of statically assigning values which are
  ```python
  [name,[name]....] = [value]
  or
  [name,[name]....] = [value,[value]....]
  ```
  in the first one all the names will be assigned to the given value and in the second one each name will have its own value due to which it will throw an error or will fail to assign if the number of variables on the left side are not equal to the ones on the left
- ### Data types
  the presently available data types are :-<br>
  - ##### Number
    A generic integer<br>declaration
    ```
    num  = 7878
    ```
  - ##### Double
    a floating point number<br>declaration
    ```
    doub = 5634.834874
    ```
  - ##### Charge
    this is an alias for boolean it can have values such as
    - true | false<br>
    declaration
    ```
    bool = true
    ```
  - ##### String
    an array of characters or a line of legible characters<br>declaration
    ```
    str = "some string"
    ```
  - ##### System-string
    now this is a string which is a shrunken down version of the variable syntax and can be used to interact directly with the language<br>declaration
    ```
    sys_str = print "hello world"
    ```
    now in this snippet if I write 
    ```
    {str}
    ``` 
    somewhere in the program later on it will become print "hello world" and the code will print hello world at that point if you print a system string then it will be printed in purple color
- ### Operators
  - #### +
    this will add two doubles or numbers and apply implicit conversions if the operands are charges then it will apply the && boolean operator and if the operand is a string then it will concatinate both sides
  - #### -
    basic arithmatic subtraction for numbers and double if the two operands are charges then it will do a boolean and with a boolean not on the second operand and if the operands are strings then it will return the higher of the two strings based on ascii codes<br>eg:-<br>
    charges
    ```python
    bool1 = true
    bool2 = false
    bool3 = bool1 - bool2 # true
    bool3 = bool2 - bool2 # false
    bool3 = bool2 - bool2 # false
    bool3 = bool1 - bool1 # false
    ```
    strings
    ```python
    str = "hzzzzz"
    joj = "aelloaaaa"
    ans = str - joj
    print "{ans}" # output will be hzzzzz
    ```
  - #### *
    basic arithmatic multiplication for numbers and Doubles,boolean and for charges and a recursive concatination for strings<br>eg:-
    ```python
    str = "hell"
    num = 3
    args = "ooo"
    ans1 = str*num
    ans2 = str*args
    print "answer1 is {ans1}"
    print "answer2 is {ans2}"
    # both will yield the same result "hellhellhell"
    ```
  - #### /
    basic arithmatic division if operands are doubles and numbers, boolean and with a boolean not on the second operand if the operands are charges and substring if the operands are strings<br>eg:-
  ```python
  str = "hello"
  num = 3
  args = "arg"
  ans1 = str/num
  ans2 = str/args
  print "answer1 is {ans1}"
  print "answer2 is {ans2}"
  # both will yield the result "he"
  ```
  - #### %
    basic arithmatic modulus in numbers,this operand is not present in doubles and charges if the operands are strings then it will return a substring of the string <br>eg:-
    ```python
    str = "hello"
    num = 3
    args = "arg"
    ans1 = str%num
    ans2 = str%args
    print "answer1 is {ans1}"
    print "answer2 is {ans2}"
    # both will yield the result "llo"
    ```
  - #### &
    boolean/logical "and" works the same way as python's "and" or java's "&&"<br>it is also bitwise "and" for number operators
  - #### |
    boolean/logical "or" works the same way as python's "or" or java's "||"<br>it is also bitwise "or" for number operands <br>This can also be used to find the string with a higher ascii value<br>eg:-
    ```python
    str1 = "hello"
    str2 = "hellp"
    str = str1|str2 # str will have hellp 
    ```
  - #### ^
    boolean/logical "xor" works the same way as java's "^"<br>it is also bitwise "xor" for number operands<br>it can also be used on strings to find the smaller of two strings<br>eg:-
    ```python
    str1 = "hello"
    str2 = "hellp"
    str = str1^str2 # str will have hello 
    ```
  - #### not
    boolean logical "not" works the same way as python's "not" and java's "!" but has a different declaration
    <br>declaration:-
    ```python
    not [var_name] [boolean value]
    ```
    example:-
    ```python
    bool1 = true
    bool2 = false
    not bool3,bool2  # bool3 will hold the value true
    not bool3,bool1  # bool3 will hold the value false
    ```
  - #### > :-
    works how you expect it to,returns charge
  - #### < :- 
    works how you expect it to,returns charge
  - #### ? :-
    this is relational "equal to" or "=="<br>eg
    ```python
    str1,str2,str3 = "hello","help","hello"
    bool = str1?str2  # false
    bool = str2?str3  # false
    bool = str3?str1  # true
    ```
  - #### ! :-
    this is not boolean "not" this is relational "not equal to" or "!=" like so<br>
    eg:-
    ```python
    str1,str2,str3 = "hello","help","hello"
    bool = str1!str2  # true
    bool = str2!str3  # true
    bool = str3!str1  # false
    ```
- ### Special Functions
  #### if else<br>
  so the structure is very simple like so<br>
  ```python
  if [boolean args] {
  	[body]
  }
  else {
  	[body]
  }
  ```
  plese note that the else block **CAN NEVER** be after the curly braces that is like this
  ```python
  if [boolean args] {
  	[body]
  } else {
  	[body]
  }
  ```
  this is not allowed and will mess the code up<br>now all boolean operators work as if's arguments but please not that there are some differences in this language then your traditional languages that is the equals to, how so? I shall demonstrate this<br>this is viper code
  ```python
  args1,args2 = "hello"
  if args1,args2 are equal {
  	[body]
  }
  ```
  this here is synonymous go code
  ```go
  args1,args2 := "hello","hello"
  if args1 == args2 {
  	[body]
  }
  ```
  as you can see the "==" is replaced with the suffixed function call " are equal" this creates the code more verbose and this will work in every situation, you can call the " are equal" function on as many operands as you want like this
  ```python
  args1,args2,args3,args4 = 20
  if args1,args2,args3,args4 are equal {
  	[body]
  }
  ```
  This is eligible code and will infact execute
  <br>As " are equal" is a method you can use syntactic sugar to make the code even more verbose like this
  ```python
  args1,args2,args3,args4 = 20
  if args1 and args2 and args3 and args4 are equal {
  	[body]
  }
  ```
  is eligible code
  <br><br>
  ####  while<br>
  will work with any charge value and operators and also supports the same boolean deciphering as "if" so the following code is eligible
  ```python
  a,b = 2,2
  while a and b are equal {
  	# do something
  }
  ```
  to do an infinite loop just write
  ```python
  while {
  	# whatever
  }
  ```
  <br>#### Input
  To take input from the user you have a simple line
  ```python
  input [var],[line to print in double quotes]
  ```
  so just like in python the line will be printed and you have to write your option in it and it will get stored in the variable "var" which you can use later
  <br><br>
  #### printin
  There are two types of print statements print and pure_print<br>
  ```python
  print[args in double quotes]
  and
  pure_print[args in double quotes]
  ```
  print is equivalent to System.out.println() and pure_print is equivalent to System.out.print()
## How to run<br>
- ### using shell script<br>
  I have added shell script in the following code for faster execution **This will work in linux or any other UNIX os which use bash script or have bash** <br>here are some lines you will have to write on the bash before actually being able to execute the shell script
  ```bash
  $ chmod +x viper  # makes the script executable
  $ export PATH=$PATH:~/Documents/VIPER-master/src/linux # you can also use the mac branch as both are the same
  ```
  after executing the given lines the following commands will become valid only for **THE PRESENT TERMINAL SESSION** 
  ```shell
  viper run [args]
  # or
  viper -r [args]
  ```
  this will execute the desired file another alternative to this is
  ```shell
  viper run -d
  # or
  viper -r -d
  ```
  this will run lethalityTest.vpr you can change the contents in it if you want to run the viper shell with the bash just type
  ```shell
  viper shell
  # or
  viper -s
  ```
  <br>For the present commands which were added to the script type 
  ```shell
  viper --help
  # or
  viper help
  # or
  viper
  ```
  <br><br>
- ### Using PowerShell
  I use linux and I have neither learned nor do I use powershell so I don't know powerShell BUT I am going to learn it moving on and when I do then this section will be more appealing then it is right now
- ### using go code<br>
  open the terminal or ide and compile and run the go code then it will ask for an argument "> " then just type in "run fileName.vpr", the file should be in the same directory as the file or pwd(present working directory), if you want quicker access then modify the lethalityTest.vpr and type "run -d" as it is set to default path <br>type quit or control-c to exit<br><br>
## About the file<br><br>
the extension for the file should be .vpr<br><br>
## The code<br>
Okay this is solely for people who are curios and want to dive deep in mu monolthic code I shall explain the core ideals and basically how my code works taking it from the easist part to the hardest
- ### Operators
  Oh this thing gave me nightmares not because it was hard but because it was boring,redundant and sooo buggy,so what I did was create a structure Data with three features :-
  - #### value
    creates non colored string representation
  - #### type
    the data type of the value
  - #### stringRep
    the colored representation of the string
  then I created a structure Operators which had 2 members:-
  - #### rep
    rune(char) representation of the operator
  - #### action
    a function to be called on the two operands
  then I created a list of Operators called oprList or operatorList and everytime compute was called the whole array was itterated to find the required operator which was then applied
- ### lexical variable Scoping
  well this was the first time I applied brainsfor the code and it turned out amazing to my surprise, so how was this made? Well the code's variables are seperated in a reversed tree data structure, let me explain a little more, a basic element of the program is a MethodData structure with the following composition:-
  ```go
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
  ```
  Do you see that pointer parent in ScopeNode,that points to the ScopeNode of the function that calls any given function thus it results in some kind of tree with globalScope(biggest scope) at the top and global scope points towards an empty ScopeNode,it is a recersed tree so the number of childnodes are always 1 and parents "n"<br>The ScopeNode and MethodData used to be very simple structures but I soon realized that I needed a ton of information about the function to carry successful execution of the code and I could not use MethodData structures as my code was deeply integrated with ScopeNode structures so I had to create the owner pointer in the ScopeNode structure
- ### Assignment and computation
  This was exhaustiong,The Idea was simple,to use linked lists to calculate the answers but my limited knowledge of pointers and creation of several complex data structures slowed me down a lot(it took me somewhat 2 hours to do this)
### About the name<br>
well the name comes from the fact that i have created a language with robust and small syntax like python<br>It was named Snek initially but some friend of mine told me it was a sloppy name<br><br>
#### future of the language<br>
[click here](https://argon-sodium-vanadium.imfast.io/snekLang.html)
