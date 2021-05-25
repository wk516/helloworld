package main

import "fmt"

func print() {
	fmt.Println("hello")
}

func name(name string) {
	fmt.Println(name)
}

func main() {
	print()
	name("wenny")
}
