// since it is belongs to diff package thats why it is kept in seperate folder --> just for organizing
package somepackage
import "fmt"

// if you want to use this function in  other packages it must be exported and to do that you just need to captalize  First letter of function name
// same rule goes for variables
func HelloWorld(){
	fmt.Print("Hello World !")

}