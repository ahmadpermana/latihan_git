package main

import "fmt"

func main() {
	if !revert() {
		hello()
	} else {
		fmt.Println("di revert")
	}
}

func revert() bool {
	return false
}

func hello() {
	fmt.Println("Hello dunia!")
}
