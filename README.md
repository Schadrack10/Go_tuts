# Golang tuts


# install Golang
-head golang.org
-head to downloads/ download your preffered type
-open terminal and type "Go" , should see files related to Golang
-on Vscode we install an extension called Go => for Go syntax support

#definition
go is a fast , statically typed, compiled language
go is a general purpose language




# advantages of Golang

# Concurrent and Parallel Programming:
One of Go's standout features is its built-in support for concurrent programming through goroutines and channels. Goroutines are lightweight threads, and channels facilitate communication between goroutines, making it easy to develop concurrent and parallel programs. This is crucial for efficiently utilizing modern multi-core processors.

# Efficient Compilation and Execution:
Go is known for its fast compilation times and the creation of statically linked binaries. The language's simplicity contributes to quick development cycles, and the resulting binaries are standalone and efficient, with no external runtime dependencies. This makes deployment straightforward and contributes to high-performance execution.

# Simplicity and Readability:
Go's syntax is intentionally simple and easy to read. It lacks some of the complexity found in other languages, which can lead to more straightforward code. The language's design encourages clean and idiomatic code, making it easier for developers to understand and maintain. This simplicity also aids in onboarding new developers to a codebase.

# Garbage collection & Dellocating memory
 memory management is handled automatically by a garbage collector. The garbage collector (GC) is responsible for identifying and reclaiming memory that is no 
 longer in use by the program, preventing memory leaks and making memory management more convenient for developers.

# creating variables and dataTypes
//strings >>

 var firstName string = "mario"  >> specifying datatype
 var firstName = "mario" >> without specifying the datatype but go is smart enough to know that it is a string
 var firstName string  >> create the string but we do not assign a value
 name := "mario" >> it is only used the first time when creating and assigning , cannot be used outside a function

//number >>  

    var age int = 30
    var age int
    age:= 30
    var age int8 = 5 >> specifying the amounts of bits and memory
    var age uint = 4 >> this ensures that we do not assign a negative number
    var age uint8 = 255 >> we can add bits to unit , since we do not add negative numbers we can go up to higher numbers
    var score float32 = 25.95 >> adding floats/numbers with decimals
    var score float64 = 25344434455454.95 >> adding floats/numbers with decimals this has a larger range of number
    score := 1.5 >> go reads and understands that this a float


# fmt > module used for formatting and printing data

1.Println() : logs data and prints a new line
eg. fmt.Println()

2.Print(): logs and prints data but does not create a new line

3.printf() : formats and prints strings
[%v >> is a format specifier] ..could be anyletter but have different functionalities
A>> fmt.printf("my age is %v and my name is %v", age , name)  >> replaces the specifier with the variables on the right
B>> fmt.printf("my age is %q and my name is %q", age , name)  >> replaces the specifier with the variables but adds quotes,this will not work on integers only strings
C>> fmt.printf("my age is %T and my name is %T", age , name)  >> replaces the specifier with the variables but gets the type of the variable instead of variable itself
D>> fmt.printf(my age is %f", age )  >> replaces the specifier with the variables this is used to output floats >> 552.55
    fmt.printf("my age is %0.1f", age )  >> when adding a number inbetween the percentage and the f this allows us to limit or control the decimal point we want eg.0,1 is one decimal point, 0.2 is two decimal point

4.Sprintf() : saves and returns the formatted strings
    var str =  fmt.Sprintf("my age is %v and my name is %v", age , name) >> string stored in variable

# Log > package used for formatting and printing data

consist of methods similar to the fmt package , we get access to format specifiers and more
note: log will prit with exact times and data of the logs [2024/01/17 00:07:49]

	log.fatal()  //print error
	log.println() // print in a new line
	log.printf()  // print  variable with format specifiers

 example:
   -2024/01/16 23:41:24 data printed


# Type conversion

 if we attempt to assign y variable it will throw an error regardless of the variales being the same types , this is because we have different integer types
 of different sizes

	 var x int32 = 1
	 var y int16 = 2
	 x = y 
  
 inorder to  have this working we need to use type converter <b>int32()</b>: the function will convert the type to an int32 hence giving them same value

	x = int32(y)


# strconv package 
allows us to convert the string to a different datatype

1.converting from string to integer using Atoi , this will return two values , the returned value and the error  val,err := strconv(variableName)

    package main
    
    import (
        "fmt"
	"strconv"
	)
	
    func main() {
        // Example string containing an integer
        str := "123"

	// Using Atoi to convert the string to an integer
	if intValue, err := strconv.Atoi(str); err == nil {
		fmt.Println("Integer Value:", intValue)
	} else {
		fmt.Println("Error:", err)
	}

	// Example with a non-integer string
	nonIntStr := "abc"
	if intValue, err := strconv.Atoi(nonIntStr); err == nil {
		fmt.Println("Integer Value:", intValue)
	} else {
		fmt.Println("Error:", err)
	}
     }
     
1.converting from integer to string using Itoa ,unlike Atoi , the Itoa function does not return multiple values , so we imidiately assign the converted value to
  a string 

	package main
	
	import (
		"fmt"
		"strconv"
	)
	
	func main() {
		// Example integer
		intValue := 123
	
		// Using Itoa to convert the integer to a string
		str := strconv.Itoa(intValue)
	
		// Printing the result
		fmt.Println("String Value:", str)
	}

 
# Constants

In golang we use Constants to let the program know our values are expressions whose values are know at compile time
basically letting the program understand the values you set up for the variables defined in the paranthesis

	const (
	  y = 4
	  z = "Hello"
	)

Iota is a function used to construct a function , basically a constructor
the iota will automatically assing to the other variable within the constant

	   type grade int
	    const (
               A grades = itota
               B
               C
	    )


# Control flow
 the order in which statemnets are executed within a program

# Hash Table
 Hash function 

  
# Arrays and slices

1.create an Array
var arr [3] int = [3]int{20, 30, 40 }  >> var =>  keyword ,arr => variable name , [3] => the amount of items that will be in , int => the data type
var ages = [3] int{20,40,44}
ages[0] >> retreive data within a slice or an array with bracket notation

//check the length of the array or a slice with the len() function
len(name_of_array)

//create a slice 

      var ages = [] int{"","",""} >> when we do not place the amount of values to be added , we create a slice automatically

//we can add into a slice but not an array , with the append function, this function returms a new slice so we need to update the value of the slice

      var names = [] int{20,30}
      names = append(names,40)
      fmt.Println(names, len(names)) // [20,30,40] , 3

//slice ranges 

      var ages = [] int{20,30,40,50}
      rangeOne := ages[1:3] >> this extract data from position 1 inclusive to position 3 exclusive
      rangeTwo := names[2:] >> this will extract everything from position 2 to the end of the slice
      rangeTwo := names[:3] >> this will extract everything from start and finish to pos 3 , it will not include position 3
      fmt.Println(rangeOne) // [30,40]  output for ages, note that we exclude position 3
      
# make() constructor

// we can use a constructor to create a slice, we can use the make() function
  this will take in two arguments :
  .1 when we have two arguments we expect the datatype and the legnth of the slice , if there are 3 args the third will specify the capacity

       sli  = make([]int,10,15)
       sli  = make(type,length,capacity)

       person := make(map[string]string) // construct a map

        

# standard Library :https://golang.org/pkg/
packages that allows us to add methods to manipulate data
# fmt, strings , sort

strings  eg. greet:= "Hello world!!"
>> Contains >> search for a word in a string eg.
>>
>>     strings.Contains(greet,"hello")
>> TrimSpace >> removes spaces around the string eg.
>>
>>     strings.TrimSpace(greet)
>> ReplaceAll >> replace a word in a string  eg.
>>
>>     strings.ReplaceAll(greet,"hello","Hii")
>> ToUpper >> converts the string to upperCase eg.
>>
>>     strings.ToUpper(greet)
>> ToLower >> converts the strings to small case eg.
>>
>>     strings.ToLower(greet)
>> Index >> retrieves the position of a word in a string , the first occurence eg.
>>
>>     strings.Index(greet,"w") //6 
>> Split >> split the word into an array, the second parameter determines how we want to spli the string into the array   eg.
>>
>>     strings.Split(greet," ")
>> 
sort eg. ages := []int{20,30,50.60,90}
>> Ints >> sorts the order of the integers in an array eg.
>>
>>     sort.Ints(ages) //this will modify the actual variable
>> SearchInts >> this will find the position of the interger in an array  eg.
>>
>>     sort.SearchInts(ages,30) // 1
>> Strings >> this method is used to sort strings alphabetically eg.
>>
>>     sort.Strings(names) // ["andy","Betty","carl"]
>> SearchStrings >> this method allos us to search for a word in an array of strings eg.
>>
>>     sort.SearchStrings()

# Loops 
>>used to iterate through an array or a slice 
 eg. method 1  

      names := []string{"Tobby","Andy","Carl","Bettany"}
      for z:=0; z<len(names); z++ {
          fmt.Println("current value is: ", names[z]) // tobby / Andy / Carl / Bettany
      }
 eg. method 2 : using range

    names := []string{"Tobby","Andy","Carl","Bettany"}

    for index, value := range names{
        fmt.Printf("the index %v and the values is %v ", index, value) // index is 0 and value is Tobby
    }

# For Range loop : excluding values 
 eg. if we do not want the value or index we replace with an underscore else we get an error

    for _,value := range names{
      fmt.Printf('the value is %v',value)
    }

//we can use keywords eg. continue , break
continue >> skips the current iteration and continues
break >> stops the iteration completely


# Booleans and conditionals

    num := 44
    if num < 44 {
    //code here
    }else if num < 55{
      //code here  
    }else{
      //code here  
    }


# Functions
keyword: func
arguments passed in need to be described 
in the paranthesis we specify the datatype of the argument/parameter the function is expecting

eg. one args type string

    func greet (n string){  
      // code here
    }

eg. multiple args
    
    func colors (a string , b [] int, c int , d func(string)){

    }

//When returnming a value we need to specify what type of value we returning 

this is after the paranthesis  and before the curly braces "int" has been added for the second time outside the paranthesis to indicate that we will be returning an interger
eg

      func sum (a int , b int) int{
        return a + b
      }

//Returning multiple values in a function 
 we add another paranthesison the right and add types we return , these are separated by commas
func add_sub (n [] int)(int, int){

    for ...{
        ...
      }
      return val1 ,val2  
    }


# Package scope
We can link two different files by specifying the same package
package main => in two files means that we create a varible/function in one file and use it in another file
note: we can only access the variable that are not function scoped 

# Maps
>> similar representation for objects in go
>> in the square bracket we specify the type of key and after we specify the type of value
>> below we have key type of string and values type of floats

eg.

        menu := map[string]float64{  
          "Soup": 4.90,
          "Pie":  4.99
        }

eg.
       
        menu := map[int]string{
          2332434:"john",
          7366738:"marry"
        }

// we can access the data in the maps
eg. meny[23324]  // john
eg. meny[23324] = "Todd" //modifies the key's value

//we can also iterate through maps, using a for range we get the key "k" and value of the keys "v"

     for i,v := range menu {
        fmt.Println(k,"-",v) // 2333 - "john"
     }

//how it is passed as a parameter in a function
func updateValue (y map[string]float64){ //code}


//We could also delete data from a map 
var mapVar =  map[string]float64{"amount":23.78,"deposited":22.32}
delete(mapVar,"deposited")

//one of the ways to check the existance of a key in a map



# By Pass value
-Go is a by-passed value language , meaning it makes copies of values when passed into functions
-When passing down a variable as a parameter in a function and attemtping to modify the variable , we only modify the copy and not the original variable
-the solution would be to return a value

# Dividing the types
groupA non Pointer types --> strings , ints , bools , arrays , floats, structs
groupB pointer wrapper values types --> function , maps , slices


# Pointers
We use the ampersin "&" to retrieve the pointer of a variable , this is usefull because we cannot modify the original variable but only the copy when it is passed in the function
We can use star "*" to retrieve the value of a pointer
eg. 
 
     name := "teddy"
     m := &name

  fmt.Println("pointer location: ",m) // pointer holding the value from name
  fmt.Println("pointer location value: ",*m) //retrieve the value of the pointer

    updatePointer(m)

//Passing pointer into function to update original data within functions
  we use the star to tell the function to expect a pointer with location addres

     func updatePointer (x *string){
       *x = "name"
     }

//note: so now what happens is we get the pointer location and store it in "m" , m originally has its own pointer location 
but the value is the pointer of the name we need to update , so we set up the parameter of the function to accept a pointer as an argument than pass it in and when the function is called we will update the copy created by the function aswell as the original variable


# Structs & custom Types
//structs are like blue prints or a schema to follow when creating data [object]
by using the type and struct keywords we can create a map that has multiple datatypes 

eg. type and struct are keywords required , we do not add comma when creating the struct , only when assigning values

    type people struct{
        name string 
        items map[string]float64
        tip float64
    }
   
    func addPeople (name string) people{
        b := people{
          name : name, //passed as a parameter
          items: map[string]float64,
          tip:0
        }
        
        return b
    }
    
// after creating  a struct assignment can vary 

1. variable : we initialise with the person type , just like we would add a datatype next to the variable eg. var p1 int
   now p1 can only contain the fields specified in the person struct , we do not use bracket but only dot notation

         type person struct {
	      name  string
	      age   int
	      addr  string
	      phone string
         }

         var p1 person
   //using dot notation to add data to person type variable


         p1.age = 20
         p1.name = "Steven"

// we cannot add fields that are not defined in the 

	 p1.race = ""  //error because we do not have race defined within the people struct

2. using a list from a struct, the type is now people , it simply replaces the type and we add people


         var people = []people{ {name:"",age:0 ,addr:"",phone:""}}

        
3. 

# Receiver functions
>> receiver functions allows us to take in data into the function inorder to create a method
  eg. the first parathensis which receives the people struct and makes it a method for the people struct
   n is a variable , people is the struct we pass in , n becomes the variable that holds the struct data

        func (n people) format() string{
           //code...
           n.items // to access the data in the struct
           return val
        } 

        people.format() //we than call the method 


# Receiver functions with Pointers
>>>  we can update struct data by passing pointer in receiver functions  , we pass the pointer in the struct itself , with receiver functions Go will understand that it needs to update the struct data and we just simply add the star on bill parameter passed as the receiver 

        func (b *bill) updateTip (tip float64){
          //(*b).tip = tip // here we de-reference the struct, but we don not need to do these go is samrt enought to know that we pass ina pointer
          b.tip = tip
        }      

//main.go 

      package main

      import "fmt"

      func main() {
        myBill := createNewBill("Shakes's bill")
        myBill.updateTip(10)
        myBill.addItems("chocolate",3.15)

        fmt.Println("created new bill: ", myBill.format())
        // fmt.Println("new bill: ", myBill)

        //  myOrder := createOrder("New christams order")
        //  fmt.Println("your order is ready: ", myOrder)
        //  fmt.Println(myOrder.formatOrder())

      }

//Bill.go

      package main

      import "fmt"

      type bill struct {
        name  string
        items map[string]float64
        tip   float64
      }

      func createNewBill(name string) bill {
        b := bill{
          name:  name,
          items: map[string]float64{"pie": 5.99, "cake": 3.99},
          tip:   0,
        }

        return b
      }

      // format bill function
      func (b bill) format() string {
        fs := "Bill Breakdown: \n"
        var total float64 = 0

        //list items
        for k, val := range b.items {
          fs += fmt.Sprintf("%-25v ...$%v \n", k+":", val)
          total += val
        }

        // add tip 
        fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)


        //total
        fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total + b.tip)

        return fs
      }

      //update
      func (b *bill) updateTip (tip float64){
        //(*b).tip = tip // here we de-reference the struct, but we don not need to do these go is samrt enought to know that we pass ina pointer
        b.tip = tip
      }  

      //add an item to the bill
      func (b *bill) addItems(name string , price float64){
        b.items[name] =  price
      }


//order.go >> my own created file for testing

     
     package main

     import "fmt"

      type order struct {
        orderId        int
        orderItems     map[string]float64
        orderReference string
      }

     func createOrder(ref string) order {
      newOrder := order{
        orderId:        123764783,
        orderItems:     map[string]float64{"Brown Bread": 23.60, "Nike shoes": 50.66},
        orderReference: ref,
      }

	     return newOrder
    }

    func (o order) formatOrder() string {

      var text string = ""

      for k, v := range o.orderItems {
        text += fmt.Sprintf("%-25v R%v \n", k, v)
      }

      return text
    }



# User Input
>>> this functionality allows golang to retrieve user input from the terminal 
>>> import "bufio" and "os" modules which will allow us to use the functionality
>>>  1. reader := bufio.NewReader(os.Stdin) //initialize the input functionality
>>>  2. input, err := reader.ReadString('\n') // after reading we get back the message and error
>>>  3. strings.TrimSpace(input) // this is removing the white space from the user input
eg.

    import (
      "bufio"
      "os"
      "strings"
    )

    func test()(string , error){
      	reader := bufio.NewReader(os.Stdin) //initialize the input functionality
        input, err := reader.ReadString('\n') // after reading we get back the message and error

        return strings.TrimSpace(input), err
    }

# Switch statements 
>> switch statement is a function that allows us to work with different possiblity based on cases 
>> similar to how we would use switch in javascript
>> 
	    	switch opt {
	          case "a":
	            fmt.Println("chose a")
	          case "t":
	            fmt.Println(tip)
	          case "s":
	            fmt.Println("Saving Bill...")
	          case "c" :
	            fmt.Println("the current bill",b)
	          default:
	            fmt.Println("That was not a valid option")
	      }

# Tagless Switch statements
>> when a switch statement does have a tage , the cases become booleans , if the confition meets than block of code within the case runs

    	    	switch  {
	          case a > 5: //example
	            fmt.Println("chose a")
	          case c == 1 :
	            fmt.Println("the current bill",b)
	          default:
	            fmt.Println("That was not a valid option")
	      }

# Parsing floats or other datatypes
>> converting stringified numbers into a float(decimal)
>> we use a package from the standard library called <b>strconv</b>
>> from the library we import the function and use the method ParseFloat
>> this method returns multipla values , onsuccess result and error

eg.

    import ("strconv")
    var numStr =  "3.44"
    floatNum,err := strcon.ParseFloat(numStr,64)

# Saving files
>> We saving files we need to save them in bytes format 
>> We need to import the os library  from the standard library
>> data is saved as a slice of bytes
>> we than call the os.writefile method ,which has 3 params , 
1. the file where we save it / followed by the name of the files
2. the data we will be saving into the file 
1. the permission code [0644]
>> the method returns an error , if there is an error we call a panic function which stops the entire proccess
>> if there are no errors we continue with the process and done we should have data crea

        func (b *bill) saveBill(){
          data := []byte(b.format())

          err := os.WriteFile("bills/"+b.name+".txt",data,0644)

          if err !=  nil {
          panic(err)
          }

          fmt.Println("Bill was saved to file")
        }

# Interfaces
>> interface groups types together based on their methods

    type shape interface{
      area() float64
      circumf() float64
    }

    type square struct{
      lenght float64
    }

    type circle struct{
      radius float64
    }


    //square methos
    func (s square) area()float64{
      return s.length * s.length
    }
    
    func (s square) circumf() float64{
      return s.lenght * 4
    }

    // circle methods
    func (c circle) area()float64{
      return  math.Pi * c.radius * c.radius
    }
    
    func (c circle) circumf() float64{
      return 2 * math.Pi * c.radius
    }

    //use a method with the interface
    func printShapeInfo(s shape){ 
      fm.Prntln("are of %T is: %0.2f \n",s,s.area())
      fm.Prntln("circumference of %T is: %0.2f \n",s,s.circumf())
    }

    func main (){
      //ready 
    }

# fmt.Scan 
The scan method allows user input , takes a pointer as an argument 
the value entered from the user will than be assinged to the pointer address passed in 


	package main
	
	import "fmt"
	
	func main() {
	    var firstNumber, secondNumber int
	
	    // Prompt the user to enter two integers
	    fmt.Print("Enter the first integer: ")
	    // Use fmt.Scan to read the first integer from the user
	    _, err1 := fmt.Scan(&firstNumber)
	
	    // Check for errors in scanning
	    if err1 != nil {
	        fmt.Println("Error:", err1)
	        return
	    }
	
	    fmt.Print("Enter the second integer: ")
	    // Use fmt.Scan to read the second integer from the user
	    _, err2 := fmt.Scan(&secondNumber)
	
	    // Check for errors in scanning
	    if err2 != nil {
	        fmt.Println("Error:", err2)
	        return
	    }
	
	    // Display the sum of the two integers
	    sum := firstNumber + secondNumber
	    fmt.Println("Sum:", sum)
	}

# Scan program 
//Program with Scan function, this will prompt user to add an integer , we than add the user input into a list and sort the list
  this will be repeated until we enter x than the program will exit

	package main
	
	import (
		"fmt"
		"sort"
		"strings"
		"strconv"
	)
	
	func PromptUser() (int,error) {
		var val string
	
		fmt.Print("Please enter an integer: ")
	
		_, errScan := fmt.Scan(&val)
	
	
		input, errConv := strconv.Atoi(val)
	
	
		if errConv != nil {
	
			if strings.ToLower(val) == "x"{
		        return 1,errConv
			}
	
			return 0,errConv
		}
	
		return input, errScan
	}
	
	func main() {
	
		const initCap = 3
		list := make([]int, 0, initCap)
	
		for {
			val, err := PromptUser()
	
	
			if err != nil {
	             
				if val == 1 {
					fmt.Println("Exit program... because you have entered x")
					break
				}
	
				fmt.Printf("Please Enter a valid integer")
	
				continue
			}
	
	
			list = append(list, val)
	
			sort.Ints(list)
	
			fmt.Println("data:", list)
		}
	}

# Protocol Packages

 In Go, the standard library includes the net package, which provides support for working with network protocols. It includes facilities for working with TCP, 
 UDP, and other network protocols. Additionally, the http package is commonly used for building web servers and clients, utilizing the HTTP protocol.
eg. net/http

# encoding/json

In Go, the encoding/json package provides functions for marshaling (converting Go data structures to JSON) and unmarshaling (converting JSON to Go data structures). The primary functions involved are json.Marshal and json.Unmarshal. Let me provide you with a brief explanation and examples for both:

# Marshal (Go to JSON):

The json.Marshal function is used to convert a Go data structure to its JSON representation.


	package main
	
	import (
		"encoding/json"
		"fmt"
	)
	
	type Person struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		City  string `json:"city"`
	}
	
	func main() {
		person := Person{Name: "John", Age: 30, City: "New York"}
	
		// Marshal the struct to JSON
		jsonData, err := json.Marshal(person)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	
		// Print the JSON data
		fmt.Println("JSON Data:", string(jsonData))
	}

 In this example, the Person struct is marshaled to JSON using json.Marshal, and the resulting JSON data is printed.

# Unmarshal (JSON to Go):

The json.Unmarshal function is used to convert JSON data into a Go data structure.

	package main
	
	import (
		"encoding/json"
		"fmt"
	)
	
	type Person struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		City  string `json:"city"`
	}
	
	func main() {
		// JSON data
		jsonData := `{"name":"Alice","age":25,"city":"London"}`
	
		// Create a Person struct
		var person Person
	
		// Unmarshal JSON data into the struct
		err := json.Unmarshal([]byte(jsonData), &person)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	
		// Print the Go data structure
		fmt.Println("Name:", person.Name)
		fmt.Println("Age:", person.Age)
		fmt.Println("City:", person.City)
	}

 In this example, JSON data is unmarshaled into a Person struct using json.Unmarshal, and the fields of the struct are printed.

Both json.Marshal and json.Unmarshal work with slices, maps, and other custom data structures as well. The json package in Go provides additional features
and options for handling JSON data, including support for custom marshaling and unmarshaling logic.



# File access

1.OS package File access

In Go, the os package provides functionalities for interacting with the operating system, including file access.
Below are some common operations related to file access using the os package:

# Opening a File:

To open a file, you can use the os.Open function. It returns a *os.File that can be used for reading or writing.

	package main
	
	import (
	    "fmt"
	    "os"
	)
	
	func main() {
	    // Open a file for reading
	    file, err := os.Open("example.txt")
	    if err != nil {
	        fmt.Println("Error opening file:", err)
	        return
	    }
	    defer file.Close()
	
	    // Read or write operations can be performed using the 'file' variable
	}

# Reading from a File:

You can use a *os.File to read data from a file. For example, you can use a bufio.Scanner to read line by line:

        package main

	import (
	    "bufio"
	    "fmt"
	    "os"
	)
	
	func main() {
	    file, err := os.Open("example.txt")
	    if err != nil {
	        fmt.Println("Error opening file:", err)
	        return
	    }
	    defer file.Close()
	
	    scanner := bufio.NewScanner(file)
	    for scanner.Scan() {
	        line := scanner.Text()
	        fmt.Println(line)
	    }
	
	    if err := scanner.Err(); err != nil {
	        fmt.Println("Error reading file:", err)
	    }
	}



# Writing to a File:

To write data to a file, you can use a *os.File with a bufio.Writer or fmt.Fprint:

	package main
	
	import (
	    "fmt"
	    "os"
	)
	
	func main() {
	    file, err := os.Create("output.txt")
	    if err != nil {
	        fmt.Println("Error creating file:", err)
	        return
	    }
	    defer file.Close()
	
	    data := "Hello, Golang!"
	    _, err = fmt.Fprint(file, data)
	    if err != nil {
	        fmt.Println("Error writing to file:", err)
	        return
	    }
	}

# Checking if a File Exists:

To check if a file exists, you can use the os.Stat function:

	package main
	
	import (
	    "fmt"
	    "os"
	)
	
	func main() {
	    _, err := os.Stat("example.txt")
	    if err != nil {
	        if os.IsNotExist(err) {
	            fmt.Println("File does not exist")
	        } else {
	            fmt.Println("Error checking file:", err)
	        }
	        return
	    }
	    fmt.Println("File exists")
	}

# Debbug golang code VScode
     
  example of directory:
  cd /Users/schadrackbotombe/Documents/Dev/Development/Server-side/golang/test-golang

     go mod init test-golang 

     go build -o

      go help build
    
