# MUSKET
code is in developement stage please wait for a release
## The Language
  Musket is an open source project which is aiming to create a formal,reliable and efficient language like go and c++ but also a language with short,concise and verbose syntax like python and ruby
<br>**sample program**
```python
hw,a,b,c = "hello, world",3,a,b

println ("hello, world")
println "hello, world"
println hw
print "hey did you know {a} = {b} = {c}"
	
x = subtract(a,b)
y = subtract a,b
z = subtract a and b

func subtract (a,b){
	return a-b
}

println " and that {x} = {y} = {z}"
{var_synt}

var_synt <- (
	println "stopping execution"
)
 ```
 output
 ```bash
 hello, world
 hello, world
 hello, world
 hey did you know 3 = 3 = 3 and that 0 = 0 = 0
 stopping execution
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
- [x] ERROR_HANDLING
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

## About the file<br><br>
the extension for the file should be .mskt<br><br>

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

