package main

import "unicode/utf8"

func main() {

	//Строки хранятся как массив байт, он неизменяемый.
	var a string = "abc"

	println(len(a))
	println(a[0])
	// a[1] = 10 //Так нельзя

	a = "абц"
	println(len(a))                    // Для хранения символов unicode нужно больше одного байта - выведет 6
	println(a[5])                      //Нежданчик. В массиве 6 байт-элементов
	println(utf8.RuneCountInString(a)) //А так посчитает символы, не байты - будет 3 (как и нужно)

	var stringFormattedVar string

	// следующие выражения равнозначны
	stringFormattedVar = "Hello,\nworld!\n\n\t\t\"quote!\""
	println(stringFormattedVar)
	stringFormattedVar = `Hello,
world!

        "quote!"`
	println(stringFormattedVar)
}
