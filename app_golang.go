package main

import "fmt"

// Моя реализация copy
func myCopy(receiver []int, origin []int) {
	for i := range receiver {
		fmt.Println(i)
		if i >= len(origin) {
			receiver[i] = 0
		} else {
			receiver[i] = origin[i]
		}
	}
}

func pyrange(start, end, step int) []int {
	// TODO: Error checking to make sure parameters are all valid,
	// else you could get divide by zero in make and other errors.

	rtn := make([]int, 0, (end-start)/step)
	for i := start; i < end; i += step {
		rtn = append(rtn, i)
	}
	return rtn
}
func main() {
	var dest []int
	dest2, dest3 := make([]int, 3), make([]int, 5)
	src := []int{1, 2, 3, 4}

	myCopy(dest, src)
	myCopy(dest2, src)
	myCopy(dest3, src)
	fmt.Println("2", dest, dest2, dest3, src) // [] [1 2 3] [1 2 3 4 0] [1 2 3 4]

	// Удаление последнего элемента слайса:
	s := []int{1, 2, 3}
	if len(s) != 0 { // защищаемся от паники
		s = s[:len(s)-1]
	}
	fmt.Println(s)

	// Удаление первого элемента слайса:
	s1 := []int{1, 2, 3}
	if len(s1) != 0 { // защищаемся от паники
		s1 = s1[1:]
	}
	fmt.Println(s1)

	// Удаление элемента слайса с индексом i:
	s3 := []int{1, 2, 3, 4, 5}
	i := 2
	if len(s3) != 0 && i < len(s3) { // защищаемся от паники
		s3 = append(s3[:i], s3[i+1:]...)
	}
	fmt.Println(s3)

	ms1 := pyrange(1, 101, 1)
	ms1 = append(ms1[:10], ms1[90:]...)
	fmt.Println(ms1)

}
