package main

import "fmt"

func foo(ziro int) int {
	// Make panic
	// panic("unexpected")
	return 7 / ziro
}

func main() {
	// выполняется при завершении main
	// defer func() {
	// 	// вызываем recover и сравниваем с nil
	// 	if r := recover(); r != nil {
	// 		fmt.Println("My error:", r) // выведет "unexpected"
	// 	}
	// }()
	foo(0) // внтури foo срабатывает паника
	fmt.Println("Вы не увидите это сообщение, так как в foo возникла паника")
}
