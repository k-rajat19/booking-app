// package to which this files belongs to--> all files which belongs to same package can access functions and variables even without using export/import concept ! 
// make sure you run all files to get the above thing works -- > go run .
package main
import "fmt"

func helperFunction(){
	fmt.Print("This is helper Function form helper.go")

}