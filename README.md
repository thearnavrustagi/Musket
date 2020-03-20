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

var something -> {
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
## About the file<br><br>
the extension for the file should be .vpr<br><br>
## How to run<br>
open the terminal or ide and compile and run the java code then it will ask for an argument "> " then just type in "run fileName.txt", the file should be in the same directory as compiler, if you want quicker access then modify the SnekTest.txt and type "runD" as it is set to default path <br>type quit or control-c to exit<br><br>
## The code<br>
I have tried an OOP approach in this compiler instead of the traditional approach of using trees(okay I had to use trees for the assignment) and I feel now that it will ease further developement as adding new features would just be a matter of adding enums and class which can or can not be easier than the tree approach it can be buggy<br><br>
### Version 0.5 (B3TA)<br>
What you see up there is the DESIRED syntax I have handled basics such as printing and assignment but the if else is under developement 
<br><br>
### About the name<br>
well the name comes from the fact that i have created a language with smaller and easier syntax than python<br><br>
#### future of the language<br>
[click here](https://argon-sodium-vanadium.imfast.io/snekLang.html)
