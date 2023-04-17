package main

import "fmt"

func sumn(n uint) uint {
	fmt.Println(n)
	if n <= 0 { // Забавно, если поставить n < 0, тогда функция уйдет в бесконечную рекурсию,
		//так как число с типом uint не может быть отрицательным, то после вычитания из нуля еденицы, число станет самым большим
		//значением для int64 (на моей системе int будет int64) - то есть, 2^64.
		return 0
	} else { // рекурсивная ветка
		return n + sumn(n-1)
	}
}

func main() {
	fmt.Println(sumn(0))
	var a uint = 0
	b := a - 1
	fmt.Println(b)
}
