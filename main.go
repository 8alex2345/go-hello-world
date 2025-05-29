package main

import (
	"fmt"
)

func main() {
<<<<<<< HEAD
	matrix := make([][]int, 10)
	counter := 0
	for x := 0; x < 10; x++ {
		matrix[x] = make([]int, 10)
		for y := 0; y < 10; y++ {
			matrix[x][y] = counter
			counter++

		}
		fmt.Println(matrix[x])
	}

=======
	fmt.Println("Hello, world!")
>>>>>>> 0ed05ddc815f190c9296c8254797ff5d6c1df042
}
