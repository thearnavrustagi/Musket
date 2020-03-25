## The code<br>
Okay this is solely for people who are curious and want to dive deep in mu monolthic code I shall explain the core ideals and basically how my code works taking it from the easist part to the hardest
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
  This was exhausting,The Idea was simple,to use linked lists to calculate the answers but my limited knowledge of pointers and creation of several complex data structures slowed me down a lot(it took me somewhat 2 hours to do this)
