package main

import "fmt"

// прим2. Глобальная константа и переменные
// нельзя объявлять нотацие :=
const program = "My application"

var name string
var ver = "v1.0.0" // прим2. инициализируем глобальную переменную

func main() {
	// прим2. изменяем глобальную переменную name
	name = "Вася"
	fmt.Println("Привет," + name + "!")
	fmt.Println("Добро пожаловать в", program, ver)

	i := 10
	if i == 10 {
		// прим.1 изменяем значение переменной i
		i += 5
		if i == 15 {
			// прим.1 в этом блоке создается новая переменная i, которая
			// перекрывает определенную выше переменную с таким же именем
			// следует избегать таких ситуаций на практике
			i := 7
			fmt.Println(i)
			// прим.1 область видимости этой переменной ограничена блоком
		}
	}
	fmt.Println(i)
}
