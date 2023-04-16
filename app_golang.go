package main

import (
	"fmt"
	"reflect"
)

func main() {

	// mySlice := make([]TypeOfelement, LenOfslice, CapOfSlice)
	// mySlice := make([]int, 0)     // слайс [], базовый массив []
	// mySlice := make([]int, 5)     // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	// mySlice := make([]int, 5, 10) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0 0 0 0 0 0]

	// 	Аргументы функции make:
	// Тип слайса (пустые квадратные скобки и тип элемента слайса).
	// Длина слайса. Если не передана, то по умолчанию равна нулю.
	// Ёмкость слайса — размер базового массива. Если значение не передано, то по умолчанию равна длине слайса.
	mySlice := make([]int, 5, 5) // слайс [0 0 0 0 0], базовый массив [0 0 0 0 0]
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	mySlice = append(mySlice, 5)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	mySlice = append(mySlice, 6, 7, 8, 9, 10)
	fmt.Println(mySlice)
	fmt.Println(mySlice[0:20])
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))

	var arr [3]int
	fmt.Println(reflect.TypeOf(arr))

	var sls []int
	fmt.Println(reflect.TypeOf(sls))

	weekTempArr := [7]int{1, 2, 3, 4, 5, 6, 7}
	workDaysSlice := weekTempArr[:5]
	weekendSlice := weekTempArr[5:]
	fromTuesdayToThursDaySlice := weekTempArr[1:4]
	weekTempSlice := weekTempArr[:]

	fmt.Println(workDaysSlice[2])
	workDaysSlice[2] = 7
	fmt.Println(workDaysSlice, len(workDaysSlice), cap(workDaysSlice))                                        // [1 2 3 4 5] 5 7
	fmt.Println(weekendSlice, len(weekendSlice), cap(weekendSlice))                                           // [6 7] 2 2
	fmt.Println(fromTuesdayToThursDaySlice, len(fromTuesdayToThursDaySlice), cap(fromTuesdayToThursDaySlice)) // [2 3 4] 3 6
	fmt.Println(weekTempSlice, len(weekTempSlice), cap(weekTempSlice))

	// Внимание: append не изменяет переданный ей слайс, а создаёт новый на основе переданного.
	// Внимание: Емкость среза - это количество элементов в базовом массиве, !считая от первого элемента в срезе! - это
	//и объясняет, почему для a в примере cap - 4, а для b -2, хотя у них один базовый массив.
	a := []int{1, 2, 3, 4}
	b := a[2:3] // b = [3]
	fmt.Println(b, len(b), cap(b))
	b[0] = 5
	b = append(b, 7)
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [3 7] 2 2
	b = append(b, 8, 9, 10)
	b[0] = 11
	fmt.Println(a, len(a), cap(a)) // [1 2 3 7] 4 4
	fmt.Println(b, len(b), cap(b)) // [11 7 8 9 10] 5 6 // [1 2 3 4 5 6 7] 7 7

	s := make([]int, 4, 7) // [0 0 0 0], len = 4 cap = 7
	// 1. Создаём слайс s с базовым массивом на 7 элементов.
	// Четыре первых элемента будут доступны в слайсе.

	slice1 := append(s[:2], 2, 3, 4)
	fmt.Println(s, slice1)
	// 2. Берём слайс из первых двух элементов s и добавляем к ним три элемента.
	// Так как суммарная длина полученного слайса (len == 5) меньше ёмкости s[:2] (cap == 7),
	// то базовый массив остаётся прежним.
	// Слайс s тоже изменился, но его длина осталась прежней

	slice2 := append(s[1:2], 7)
	fmt.Println(s, slice1, slice2) // [0 0 7 3] [0 0 7 3 4] [0 7]
	// 3. Здесь также базовый массив остаётся прежним, изменились все три слайса

	slice3 := append(s, slice1[1:]...)
	fmt.Println(len(slice3), cap(slice3)) // 8 14
	// 4. Длина s и slice1[1:] равна 4, длина нового слайса будет равна 8,
	// что больше ёмкости базового массива.
	// Будет создан новый базовый массив ёмкостью 14,
	// ёмкость нового базового массива подбирается автоматически
	// и зависит от текущего размера и количества добавленных элементов

	// 5. Легко проверить, что slice3 ссылается на новый базовый массив
	s[1] = 99
	fmt.Println(s, slice1, slice2, slice3)
	// [0 99 7 3] [0 99 7 3 4] [99 7] [0 0 7 3 0 7 3 4]

	// В примере выше, не очень понятно, будут ли новые слайсы ссылаться на
	//тот же базовый массив или отправят свои копии в новый массив.
	//Поэтому на практике функцию append рекомендуют лишь для присвоения
	//слайса самому себе: s = append(s, b).
}
