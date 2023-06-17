package tools

import "fmt"

func EchoSuccess(str string) {
	fmt.Printf("\033[1;32;40m%s\033[0m\n", str)
}

func EchoError(str string) {
	fmt.Printf("\033[1;31;40m%s\033[0m\n", str)
}
