package main

import "fmt"

const id = 100

func main() {
	// пример присвоения переменным разных типов значения нетипизированной константы
	// если константу id типизировать явно, или заменить на переменную, то код ниже будет вызывать ошибку.
	var i int64 = id
	var f float64 = id
	fmt.Println("i=", i, "f=", f)

	var a float64
	a = 5.0 + 5
	fmt.Println(a)

	const pi = 3.14159
	const doublePi = pi * 2
	const version = "1.0.0"

	// // эквивалентно
	// const (
	// 	pi       = 3.14159
	// 	doublePi = pi * 2
	// 	version  = "1.0.0"
	// )
	fmt.Println(version, pi, doublePi)

	// Нетипизированные константы
	const intConst = 5
	const floatConst = 5.0
	const runeConst = 'A'
	const strConst = "Hello, world!"
	const boolConst = true
	fmt.Println(intConst, floatConst, runeConst, strConst, boolConst)

	// Если в группе у константы не указано значение, то оно равно значению предыдущей константы.
	const (
		pi_ = 3.1415
		e
		name = "John Doe"
		fullName
	)
	fmt.Println("pi =", pi, "e =", e)
	fmt.Println("name =", name, "fullName =", fullName)

	// Если при объявлении вы указываете тип константы явным образом, она становится
	// типизированной и подчиняется правилам сильной типизации Go
	// const flag uint8 = 128
	// var i int = flag
	// fmt.Println(i)
}
