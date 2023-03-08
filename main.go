package main

import (
	"fmt"
	"learn-actions/maths"
)

func main() {
	sum := maths.AddNumbers(2, 3)
	diff := maths.SubtractNumbers(3, 2)

	fmt.Printf("2 + 3 = %d", sum)
	fmt.Printf("3 - 2 = %d", diff)
}
