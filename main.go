package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	if num >= 90 && num <= 100 {
		fmt.Println("Отлично")
	} else if num > 70 && num < 90 {
		fmt.Println("Хорошо")
	} else if num >= 50 && num < 70 {
		fmt.Println("Удовлетворительно")
	} else if num < 50 {
		fmt.Println("Неудовлетворительно")
	}
}
