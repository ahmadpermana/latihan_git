package main

import "fmt"

func main() {
	if !revert() {
		hello()
	}
}

func revert() bool {
	return false
}

func hello() {
	fmt.Println("Hello Worlds!")
	fmt.Println("Hello dunia!")
}
