package main

import "fmt"

func print() {
	fmt.Println("hello")
}

func name(name string) {
	fmt.Println(name)
}

func menu() {
	fmt.Println("1. Enter your name")
}

func main() {
	menu()
	print()
	name("wenny")
}
