package main

import "fmt"

func main() {

	string := "a string value"

	stringPointer := &string

	fmt.Printf("string:%p", &string)
	fmt.Println()
	string = "string changed"
	fmt.Printf("string:%p", &string)

	fmt.Println(string)
	fmt.Println(stringPointer)
	fmt.Println(*stringPointer)

}
