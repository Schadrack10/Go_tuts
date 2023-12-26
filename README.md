# Go_tuts
Golang tutorial

# install Golang
-head golang.org
-head to downloads/ download your preffered type
-open terminal and type "Go" , should see files related to Golang
-on Vscode we install an extension called Go => for Go syntax support

#definition
go is a fast , statically typed, compiled language
go is a general purpose language




# advantages of Golang
1. fast


#creating variables and dataTypes
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


# fmt > module used for formatting and printing strings

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


# standard Library :https://golang.org/pkg/
packages that allows us to add methods to manipulate data
eg.fmt, strings , sort

strings  eg. greet:= "Hello world!!"
>> Contains >> search for a word in a string eg. strings.Contains(greet,"hello")
>> ReplaceAll >> replace a word in a string  eg. strings.ReplaceAll(greet,"hello","Hii")
>> ToUpper >> converts the string to upperCase eg. strings.ToUpper(greet)
>> ToLower >> converts the strings to small case eg. strings.ToLower(greet)
>> Index >> retrieves the position of a word in a string , the first occurence eg. strings.Index(greet,"w") //6 
>> Split >> split the word into an array, the second parameter determines how we want to spli the string into the array   eg. strings.Split(greet," ") 

sort eg. ages := []int{20,30,50.60,90}
>> Ints >> sorts the order of the integers in an array eg.sort.Ints(ages) //this will modify the actual variable
>> SearchInts >> this will find the position of the interger in an array  eg. sort.SearchInts(ages,30) // 1
>> Strings >> this method is used to sort strings alphabetically eg. sort.Strings(names) // ["andy","Betty","carl"]
>> SearchStrings >> this method allos us to search for a word in an array of strings eg.sort.SearchStrings()

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


# functions
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
