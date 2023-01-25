package main

//Задание 2. Шахматная доска

import (
	"fmt"
)

func main() {
	var width int
	var height int
	fmt.Println("Введите ширину:")
	fmt.Scan(&width)
	fmt.Println("Введите высоту:")
	fmt.Scan(&height)

	for i := 0; i != height; i++ {
		fmt.Println()
		for j := 0; j != width; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
	}

}
