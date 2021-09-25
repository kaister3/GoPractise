package main

import (
	"awesomeProject/grammerLearn/util"
	"fmt"
)

func main() {
	a := 5
	b := 10

	Change(a, &b)

	fmt.Println(a, b)

	fmt.Println(FormatString("kris", 23, 177))

	util.Lengthen(1, 2, 3, 4, 5, 6, 7, 8)
}

func FormatString(name string, age, height int) string {
	return fmt.Sprintf("hello %s age = %d height = %d", name, age, height)
}

func Change(a int, b *int) {
	a = 0
	*b = 12
}
