package main

import "fmt"

func Say(animal string) (v string) {
	switch animal {
	default:
		v = "heh"
	case "dog":
		v = "gav"
	case "cat":
		v = "myau"
	case "cow":
		v = "mu"
	}
	return
}

func Do(say bool) func(string) string {
	if say {
		return Say
	}
	return func(s string) string { return s }
}

func Print(who string, how func(string) string) {
	fmt.Println(how(who))
}

func main() {
	Print("cat", Say)

	f := func(s string) string { return s }
	fmt.Println(f("mmmmmmm"))

	Print("dog", Do(true))
}
