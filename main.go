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
	fmt.Println("2. Wave to me")
	fmt.Println("3. Wave back to you")

}

func main() {
	menu()
	print()
	name("wenny")
	name(:siew tuck meng")
}
