package main

import "fmt"

func main() {

	a := 1000
	b := 1000.0
	c := 1_000          //Можно разделять части числа символом "_" для удобства восприятия
	d := 01750          // восьмиричное представление, начинаенся с 0
	e := 0x3e8          // шестнадцатиричное представление
	f := 0b001111101000 // бинарное представление

	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", b)
	fmt.Printf("%d\n", c)
	fmt.Printf("0%o\n", d)
	fmt.Printf("Ox%x\n", e)
	fmt.Printf("Ob%b\n", f)
}
