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

var something <- {
	print "hello"
	var = {yulo}
}

method subtract = (a,b){
	return a-b
}

method add = (a,b){
	return a+b
}
 ```
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
 - [ ] SYNTAX<br>
   - [x] normal assignment
   - [x] forced assignment
   - [X] printing                      
   - [ ] user input                    
   - [ ] type variables                
   - [ ] if/else                       
   - [ ] while/for                      
- [ ] OPERATORS<br>
  - [ ] arithmatic                    
  - [ ] logical                       
  - [ ] relational                    
  - [ ] bitwise                       
- [ ] PROGRAMMING FEATURES<br>
  - [x] Lexical Variable Scoping
  - [x] Procedural programming        
  - [ ] Object Oriented Programming   
  - [x] Functional Programming   
  - [x] Variable Syntax
  - [ ] Smart Interpreter
  - [x] Auto Main
- [x] ERROR_HANDELING
## Unique Features and Syntax
Well the one which I just developed is called var syntax
- ### VAR_SYNTAX<br>
  well this is basically that you assign a variable like this
  ```python
  var [var_name] <- {
  	# body
  }
  ```
  and this should be outside a method<br>
  to call this you just have to write its name in the main method like so
  ```python
  [var_name]
  ```
  on being called the place where the name was will be changed by its body now why I say changes well let me give you an example
  ```python
  method main = (){
  	var a,b = 20,30
  	add
  	print "the sum of a and b is {c}"
  }
  var add <- {
  	var c = a + b 
  }
  ```
  yes this is eligible syntax but will not work right now well because I have not handled the calculation part (yes the language is even more useless then before but well calculation implementations are easy and redundant)
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
  method printf = (name1,name2){
  	pure_print("{name1} {name2}")
  }
  ```
  or
  ```python
  method main = (){
  	print "hello world"
  	printf "hello" and "world"	
  }
  method printf = (name1,name2){
  	pure_print("{name1} {name2}")
  }
  ```
  both are eligible syntax
- ### Function Calling
  This is another thing that I feel is beautiful about the language why so? well here you go 
  ```python
  method main = (){
  	var a,b = 20,30
  	println "hello,world",num
	var x = add a and b
	var c = (choose a or b)
  }
  method println = (str,num){
  	print "{str}{num}"
  }
  method add = (arg1,arg2){
  	return (arg+var)
  }
  method choose = (arg1,arg2)(
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
- ### Scoping
  Scoping is done by default in the program by that I basically mean that the values of one method don't go jumping in that of another method, you can also control it with scoping blocks an example is
  ```python
  method main = (){
  	str = "hello"
	
	scope {
		str = "bye"
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
- ### Input and printing
  So to take you have a simple line
  ```python
  input [var],[line to print in double quotes]
  ```
  so just like in python the line will be printed and you have to write your option in it and it will get stored in the variable "var" which you can use later
  <br><br>
  Now there are two types of print statements print and pure_print<br>
  ```python
  print[args in double quotes]
  and
  pure_print[args in double quotes]
  ```
  print is equivalent to System.out.println() and pure_print is equivalent to System.out.print()
## How to run<br>
- ### using shell script<br>
  I have added shell script in the following code for faster execution **This will work in linux or any other UNIX os which use bash script** so the code is executed like so 
  ```shell
  ./viper.sh run [args]
  # or
  ./viper.sh -r [args]
  ```
  this will execute the desired file another alternative to this is
  ```shell
  ./viper.sh run -d
  # or
  ./viper.sh -r -d
  ```
  this will run lethalityTest.vpr you can change the contents in it if you want to run the viper shell with the bash just type
  ```shell
  ./viper.sh shell
  # or
  ./viper.sh -s
  ```
  with **ABSOLUTELY NO SPACES**<br>For the present commands which were added to the script type 
  ```shell
  ./viper.sh --help
  # or
  ./viper.sh help
  # or
  ./viper.sh
  ```
  <br><br>
  #### important notes<br>
  - ##### It is possible that the viper.sh file is not executable when you install it, if that is the case then there are two alternatives
    - MAKE IT EXECUTABLE
      with the command
      ```shell
      chmod +x viper.sh 
      ```
    - RUN USING "sh" command
      ```shell
      sh viper.sh [args]
      ```
  - ##### second thing is that the main BINARY FILE or EXECUTABLE file should always be in the same directory
- ### using go code<br>
  open the terminal or ide and compile and run the go code then it will ask for an argument "> " then just type in "run fileName.vpr", the file should be in the same directory as the file or pwd(present working directory), if you want quicker access then modify the lethalityTest.vpr and type "run -d" as it is set to default path <br>type quit or control-c to exit<br><br>
## About the file<br><br>
the extension for the file should be .vpr<br><br>
## The code<br>
Very simple clunky POP approach (believe me OOP is ammazing this POP almost killed me) but it has scope for improvement because after this I have to arrange it into packages then it will have amazing scalability
### About the name<br>
well the name comes from the fact that i have created a language with robust and small syntax like python<br>It was named Snek initially but some friend of mine told me it was a sloppy name<br><br>
#### future of the language<br>
[click here](https://argon-sodium-vanadium.imfast.io/snekLang.html)
