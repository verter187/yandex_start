package main

import (
	"fmt"
	"reflect"
)

type Name string
type Fruit string

func main() {
	var fruit Fruit
	var name Name

	fruit = "Apple"

	// name = fruit // Ошибка типизации, хотя оба типа созданы из типа string
	name = Name(fruit) //а так, после приведения типов, работает.

	fmt.Println(name)
	fmt.Println(reflect.TypeOf(fruit))
	fmt.Println(reflect.TypeOf(name))
}
