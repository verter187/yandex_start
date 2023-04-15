package main

import "fmt"

func main() {
	// Метку можно указать для операторов:
	// break;
	// continue;
	// goto (безусловный оператор перехода, позволяет перейти в любое место кода).

	// outerLoopLabel:
	// 	for i := 0; i < 5; i++ {
	// 		for j := 0; j < 5; j++ {
	// 			fmt.Printf("[%d,%d]\n", i, j)
	// 			break outerLoopLabel
	// 		}
	// 	}
	// 	fmt.Println("End")
	// }

outerLoopLabel:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("[%d,%d]\n", i, j)
			continue outerLoopLabel
		}
	}
	fmt.Println("End")

	group := 0
	for i := 0; i < 20; i++ {
		switch {
		case i%2 == 0:
			if i%10 == 0 {
				group++
				break // break относится к ближайшему switch
			}
			fmt.Printf("%02d: %d\n", group, i)
		default:
		}
	}
}
