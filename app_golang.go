package main

import "fmt"

func main() {

	var a int
	var b string
	//Присвоение в строку переменным разных типов
	a, b = 5, "test"

	fmt.Println(a, b)

	// Реализация swap и неявная типизация
	var d, c = 5, 10
	println(d, c)

	c, d = d, c
	println(d, c)

	a *= 2
	fmt.Println("a:", a)
	a++
	fmt.Println("a:", a)

}
