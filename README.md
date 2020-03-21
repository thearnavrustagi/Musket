# VIPER
## VIPER SYNTAX (or general syntax)<br>
code is in developement stage please wait for a release
 The language is as follows
```python
__________VIPER_____________
method main = (){ # comment
	var x = 20
	var a,b,c = 10,15,17
	x = 30

	something

	print "okay a is {a}|{b}|{c}|{x}\n" # the \n is compulsory
	add <- (a,b)
}

var something <- {
	print "hello"
	var = {subtract}
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
	System.out.println("okay a is "+a+"|"+b+"|"+c+"|"+x)
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
   - [ ] while                         
- [ ] OPERATORS<br>
  - [ ] arithmatic                    
  - [ ] logical                       
  - [ ] relational                    
  - [ ] bitwise                       
- [ ] PROGRAMMING FEATURES<br>
  - [ ] Procedural programming        
  - [ ] Object Oriented Programming   
  - [x] Functional Programming   
  - [x] Variable Syntax
  - [ ] Smart Interpreter
- [x] ERROR_HANDELING
## Unique Features !!!
Well the one which I just developed is called var syntax
- **VAR_SYNTAX**<br>
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
## How to run<br>
- **using shell script**<br>
  I have added shell script in the following code for faster execution **This will work in linux or any other UNIX os which use bash script** so the code is executed like so 
  ```shell
  ./viper.sh run [args]
  ```
  this will execute the desired file another alternative to this is
  ```shell
  ./viper.sh run -d
  ```
  this will run lethalityTest.vpr you can change the contents in it if you want to run the viper shell with the bash just type
  ```shell
  ./viper.sh
  ```
  with **ABSOLUTELY NO SPACES**<br><br>
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
- **using go code**<br>
  open the terminal or ide and compile and run the go code then it will ask for an argument "> " then just type in "run fileName.vpr", the file should be in the same directory as the file or pwd(present working directory), if you want quicker access then modify the lethalityTest.vpr and type "run -d" as it is set to default path <br>type quit or control-c to exit<br><br>
## About the file<br><br>
the extension for the file should be .vpr<br><br>
## The code<br>
Very simple clunky POP approach (believe me OOP is ammazing this POP almost killed me) but it has scope for improvement because after this I have to arrange it into packages then it will have amazing scalability
### Version 0.5 (B3TA)<br>
What you see up there is the DESIRED syntax I have handled basics such as printing and assignment but the if else is under developement 
<br><br>
### About the name<br>
well the name comes from the fact that i have created a language with robust and small syntax like python<br>It was named Snek initially but some friend of mine told me it was a sloppy name<br><br>
#### future of the language<br>
[click here](https://argon-sodium-vanadium.imfast.io/snekLang.html)
