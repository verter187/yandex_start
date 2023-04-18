package main

import "fmt"

var Global = 5

func editGlobal() {

	defer func(checkout int) { Global = checkout }(Global)
	Global = 42
	fmt.Println("Cur global value:", Global)
}

func main() {
	fmt.Println("Global value:", Global)
	editGlobal()
	fmt.Println("Global value:", Global)
}
