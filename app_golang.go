package main

import "fmt"

func main() {
	thisWeekTemp := [7]int{-3, 5, 7}
	// a := thisWeekTemp[3+4] //Компилятор выдаст ошибку - выход за пределы массива
	fmt.Println(thisWeekTemp)

	// Количество элементов в массиве может быть выведено автоматически по длине списка инициализации
	rgbColor := [...]uint8{255, 255, 128}     // [255 255 128] len = 3
	rgbColor1 := [...]uint8{255, 255, 128, 1} // [255 255 128 1] len = 4
	fmt.Println(rgbColor, rgbColor1)

	// в списке инициализации можно указать только нужные элементы и их индексы
	thisWeekTemp1 := [7]int{6: 11, 2: 3} // [0 0 3 0 0 0 11]
	fmt.Println(thisWeekTemp1)

	// Размер массива может быть получен встроенной функцией
	//Размер массива известен на этапе компиляции, поэтому вычисление
	//этой функции при компиляции подменяется конкретным значением.
	fmt.Println(len(thisWeekTemp1))

	var thisMonthTemp [4][7]int // массив из четырёх недель, каждая из которых — массив из семи дней
	// b := thisMonthTemp[4][7] //Компилятор выдаст ошибку - выход за пределы массива
	fmt.Println(thisMonthTemp)

	// Многомерные массивы удобно представлять себе как массив массивов
	var rgbImage [1080][1920][3]uint8 // изображение — это массив из 1080 строк длиной в 1920 пикселей. Каждый пиксель — массив из трёх байт
	// 1080 — размер массива
	// [1920][3]uint8 — тип элемента
	// fmt.Println(rgbImage)
	// line := rgbImage[2]      // 3-я строка в изображении
	pixel := rgbImage[2][3]  // 4-й пиксель в третьей строке изображения
	red := rgbImage[2][3][1] // значение синей компоненты (второй байт) 4-го пикселя в третьей строке изображения
	// fmt.Println(line)
	fmt.Println(pixel)
	fmt.Println(red)

	var weekTemp = [7]int{5, 4, 6, 8, 11, 9, 5}
	sumTemp := 0
	lenWeek := len(weekTemp)
	// for i := 0; i < lenWeek; i++ {
	// 	sumTemp += weekTemp[i]
	// }

	for _, temp := range weekTemp {
		sumTemp += temp
	}
	average := sumTemp / lenWeek
	fmt.Println(sumTemp, lenWeek, average)
	for i := range weekTemp { // если значение элемента не используется, можно опустить вторую переменную
		// temp = 0 //Не изменит значение в массиве
		weekTemp[i] = 0 //Так изменит.
	}
	fmt.Println(weekTemp)
}
