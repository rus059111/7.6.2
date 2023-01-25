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

	count := 4

	for i := 0; i < count; i++ {
		fmt.Print("*")
		fmt.Println(" ")
		for i := 0; i < count; i++ {
			fmt.Print("*")
			fmt.Println(" ")
		}
	}

}
