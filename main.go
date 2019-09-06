package main

import "fmt"

func main() {
	hello()
	if revert() {
		fmt.Println("di revert")
	}
}

func revert() bool {
	return false
}

func hello() {
	fmt.Println("Hello dunia!")
}
