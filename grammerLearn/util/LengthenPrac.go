package util

import "fmt"

type Options struct {
	par1 int
	par2 string
}

func Lengthen(a, b int, args ...int) {
	fmt.Printf("%v", args)
}

func F1(s ...string) {
	F2(s...)
	F3(s)
}

func F2(s ...string) {
	//
}

func F3(s []string) {
	//
}

//但是如果变长参数的类型并不是都相同的呢
//使用 3 个参数来进行传递并不是很明智的选择

func F4() {
	// use struct
	F5(1, "asd", Options{})
	F5(1, "asd", Options{1, "asd"})

	// use empty interface
	F6(1, "asd", 1, "asd")
}

func F5(a int, b string, options Options) {
	//
}

func F6(values ...interface{}) {

}
