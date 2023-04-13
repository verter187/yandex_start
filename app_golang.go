package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	var r = 'Щ'
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(string(r))

	//объявить сразу блок переменных
	var (
		height, length int
		weight         float64
		name           string
		company        = "Рога и копыта"
	)
	fmt.Println(height, length, weight, name, company)

	// Операция приведения типа помогает использовать короткую нотацию
	int64Var := int64(5)
	float32Var := float32(101.3)

	// эквивалентно

	var int64Var1 int64 = 5
	var floatVar float32 = 101.3
	fmt.Println(int64Var, int64Var1, floatVar, float32Var)

	pi, e := 3.1415, 2.7183
	// при уточнении значений нельзя использовать :=, так как
	// обе переменных уже определены
	pi, e = 3.14159, 2.71828

	f, err := os.Open("myfile.txt")
	fmt.Println(f, err, pi, e)

	ver, id, pi := "v0.0.1", 0, 3.1415

	fmt.Println("ver =", ver, "id =", id, "pi =", pi)
}
