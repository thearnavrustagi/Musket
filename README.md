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
	print "okay a is {a}|{b}|{c}|{x}\n" # the \n is compulsory
	add <- (a,b)
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
  - [ ] Variable Syntax
  - [ ] Smart Interpreter
- [x] ERROR_HANDELING
 
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
