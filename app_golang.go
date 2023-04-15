package main

import "fmt"

func main() {
	// var a int = 5
	// p := &a

	// fmt.Println(a, p) //a=5 p=0xc0000b2008

	// С константами сложнее — у них забрать адрес не получится.
	//		  const c = 5
	//	  p1 := &"abc" // ошибка компиляции
	//	  p2 := &с // ошибка компиляции

	// Тип переменной, на которую создаётся указатель, должен соответствовать типу указателя.
	// var p *int
	// var a int = 5
	// var b string = "abc"
	// p = &a
	// p = &b // ошибка компиляции

	// // Литералы композитных типов создают в памяти переменную соответствующего типа
	// type A struct {
	// 	IntField int
	// }
	// // Литерал А{} создаёт в памяти переменную типа А. Затем от неё берётся указатель
	// z := &A{
	// 	IntField: 10,
	// }
	// fmt.Println(z)
	// // А ещё в Go есть встроенная функция new()
	// d := new(A) //  то же самое, что и &A{}
	// fmt.Println(d)

	// i := 42
	// r := &i
	// fmt.Println(*r) // читаем значение переменной i через указатель p
	// *r = 21         // записываем в переменную i значение 21 через указатель p
	// fmt.Println(*r)

	// m := []int{1, 2, 3}
	// b := m
	// fmt.Println(b)
	// b[1] = 4
	// fmt.Println(b)
	// fmt.Println(m)

	// incrementCopy := func(i int) {
	// 	i++
	// }

	// increment := func(i *int) {
	// 	(*i)++
	// }

	// i := 42

	// incrementCopy(i)
	// fmt.Println(i) // 42

	// increment(&i)
	// fmt.Println(i) // 43

	a := 1
	p := &a
	b := &p

	*p = 3
	**b = 5

	a += 4 + *p + **b

	fmt.Printf("%d", *p)
}
