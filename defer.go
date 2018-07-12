package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("Hello")
	fmt.Println("with")
	fmt.Println("my friend")
	fmt.Println("and")
	fmt.Println("the")
}
