package main

import (
	"runtime"
)

var a string

func main() {
	print("you are currently using ")
	if runtime.GOOS == "Linux" {
		print("Linux")
	} else {
		print(runtime.GOOS)
	}
}
