package main

import "fmt"

var jwtToken = "klasjfklsadjfsldkfjsdklfjsdalkfjasklfjkjeirju2845u23848jdksfjksdaur29874"
var code = "sayHi"

func main() {
	var username string = "Aziz"
	fmt.Println(username)
	fmt.Printf("The type of this variable is %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of this variable is %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("The type of this variable is %T \n", smallVal)

	var smallFloat float32 = 154.25
	fmt.Println(smallFloat)
	fmt.Printf("The type of this variable is %T \n", smallFloat)

	var defaultInt int // zero
	fmt.Println(defaultInt)
	fmt.Printf("The type of this variable is %T \n", defaultInt)

	var defaultString string // an empty string
	fmt.Println(defaultString)
	fmt.Printf("The type of this variable is %T \n", defaultString)

	if code != "hi" {
		myName := "Ahmed Aziz Abbassi"
		var myAge uint8 = 20
		fmt.Println(sayHi(myName))
		fmt.Printf("My name is %s and I am %d years old! \n", myName, myAge)
		fmt.Printf("My name is of type %T and my age is of type %T \n", myName, myAge)
	}
}

func sayHi(someone string) string {
	return "Hi " + someone
}
