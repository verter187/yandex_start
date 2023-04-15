package main

import "fmt"

// Реализация ENUM
// const (
//
//	Black = "black"
//	Gray  = "gray"
//	White = "white"
//
// )
const (
	Black = iota
	Gray
	White
)

// счётчик обнуляется
const (
	Yellow = iota
	Red
	Green = iota // это присваивание не обнулит iota
	Blue
)

func main() {
	// fmt.Println(Black != Gray) // true
	fmt.Println(Black, Gray, White)
	fmt.Println(Yellow, Red, Green, Blue)

	const (
		_ = iota * 10 // обратите внимание, что можно пропускать константы
		ten
		twenty
		thirty
	)

	const (
		hello = "Hello, world!" // iota равна 0
		one   = 1               // iota равна 1

		black = iota // iota равна 2
		gray
	)

	fmt.Println(ten, twenty, thirty)
	fmt.Println(hello, one, black, gray)
}
