package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to our App"
	fmt.Println(welcome)

	buffer := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our pizza : ")

	// comma ok || err ok
	input, _ := buffer.ReadString('\n')
	fmt.Println("Thanks for rating us! Your rating was", input)
}
