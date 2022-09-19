// intiliaze project - go mod init booking-app
package main
import "fmt"

//  to import a package made by us not a built-in one
import "booking-app/somepackage"
// booking-app is a  module in go.mod file



//package level varibles --> can be acessed by main and other func as well
// packageLevelVariableName :="something" ---> can't use this syntax here
// var packageLevelVariableName= "something"

func main(){
	var name="Stranger Things"
	//also name:="Stranger Things" --> not works with const ,also  you can't define its data types

	// if we don't assign variable value at the time we are declaring it go throws error and ask to define its data type
	//var somevariable

	//var somevarible string --> it doesn't throw error

	fmt.Printf("type of name is %T" , name)
	fmt.Printf("value of name is %v",name)
	fmt.Print("Hello World ",name)

	var inputUserName string
	//& --> is a pointer that references to the memory address of inputUserName variable
	fmt.Scan(&inputUserName)

	//var arrayNames= [10]string{"mike","steve","jonathan"}
	var arrayNames [10]string
	// arrayNames[0]={"mike"}
	// in go arrays are of  fixed size and you need to define its size ,also you can't mix type of data stored in an array

	// using slices --> slices doesn't need to be of a fixed size unlike in array , it can resize autom. when needed
	//arrayNames= append(arrayNames, valueof element thats needs to be appended/added in arrayNames)

     
	// for{fmt.Print("it runs infinte time")}

	for index,names:= range arrayNames{
		fmt.Print(index);
		fmt.Print(names);
	}
	// you can also add condition in for loop like this:
	// for something>0{

	// }

   //  in above example if we don't want to use index variable in our logic we can use _ in place of it because some variable has to be defined  in place of index
  /// to ignore a variable you don't want to use but it have to be declared for some  technical reasone , use _ (blank indetifier)


	  // `break` keyword inside for loop breaks the loop and exit
	  // continue keywords breaks loop but does not exit loop instead it starts next iteration

	  // if condition{}
	  testFunction("this is the string");

	   varibleToStroreIntFromTestFunction , toStoreString  :=testFunction("")
	  fmt.Print(varibleToStroreIntFromTestFunction,toStoreString)

	  // coming from helper.go (main package)
	  helperFunction();

	  // coming from  somepackage.go (somepackage package)
	  somepackage.HelloWorld();
}



func testFunction(param string)(int ,string){
	fmt.Print("hello this is test function");

	// yes you can return multiple values
	return 3,"a string"
}
