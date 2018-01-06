package main

import "fmt"

func main() {
	hello1()
	hello2()
	hello3("asdf")
}
func hello1() string {
	fmt.Println("Hello Go 1!")
	return "Hello Go 1!"
}

func hello2() string {
	fmt.Println("Hello Go 2!")
	return "Hello Go 2!"
}

func hello3(input string) string {
	output := "Hello "
	output += input
	output += "!"
	fmt.Println(output)
	return output
}

type animal struct {
	legs int
	eyes int
	hair bool
	name string
}

type mammal interface {
}

type amphibian interface {
}
