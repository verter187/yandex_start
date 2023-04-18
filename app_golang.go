package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func EvaluationOrder() {
	defer fmt.Println("deferred")
	fmt.Println("evaluated")
}

func unintuitive() (value string) {
	defer func() { value = "На самом деле" }() // круглые скобки в конце означают, что функция вызывается
	return "Казалось бы"
}

func intuitive() string {
	value := "Казалось бы"
	defer func() { value = "На самом деле" }()
	return value
}

func metricTime(start time.Time) {
	// функция Now() возвращает текущее время, а функция Sub возвращает разницу между двумя временными метками
	fmt.Println(time.Now().Sub(start))
}

func VeryLongTimeFunction() {
	defer metricTime(time.Now()) // передаём в функцию metricTime значение текущего времени и откладываем её вызов до возврата
	// Какие-то долгие вычисления
	time.Sleep(3 * time.Second)
}

func PanicingFunc() {
	defer func() {
		if r := recover(); r != nil { // встроенная функция recover останавливает панику и возвращает описание произошедшего
			fmt.Println("Panic is caught", r)
		}
	}()
	///
	///

	panic("Мне здесь совсем ничего не нравится!")
	// встроенная функция panic () вызывает панику у функции.
	// в качестве аргумента ей принято передавать причину паники. Именно она будет возвращена функцией recover

}

func main() {
	// EvaluationOrder()
	fmt.Println("Hello")
	for i := 1; i <= 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("World")
	test := unintuitive()
	fmt.Println(test)

	test1 := intuitive()
	fmt.Println(test1)

	// Операнды отложенной функции вычисляются при выполнении оператора defer, а не
	// при выполнении отложенной функции
	a := "some text"
	defer func(s string) {
		fmt.Println(s)
	}(a)
	a = "another text"

	// Оператор defer чаще всего можно увидеть с парными функциями
	// Open()/Close(), Lock()/Unlock(). Его ставят сразу после захвата ресурса, чтобы точно не забыть.
	// открываем файл
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем закрыть файл
	defer file.Close()
	// работаем с файлом
	_, err = file.WriteString("")
	if err != nil {
		log.Fatal(err)
	}

	VeryLongTimeFunction()

	PanicingFunc()
}
