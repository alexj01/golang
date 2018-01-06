package main

import (
	"./greeting"
)

func main() {

	// var message string
	// var message2 string = "AHHHHHH!"
	// var a, b, c int
	// var d, e, f int = 1, 2, 3
	// var inferredType1, inferredType2, inferredType3 = 1, false, "hello"

	salutations := greeting.Salutations{
		{"Bob", "Hello"},
		{"Mary", "Hola"},
		{"Joe", "Hey"},
		{"Alex", "Whoa"},
	}

	// add to slice
	//salutations = append(salutations, greeting.SalutationStruct{"Jimmy", "Ah"})
	//fmt.Println(salutations)

	// delete from slice
	//salutations = append(salutations[:1], salutations[2:]...)
	//fmt.Println(salutations)

	// Inside a function only for this format
	// declaredAndInitialized := 234.123
	// message = "Hello Go World!"
	// fmt.Println(message, message2, inferredType1, inferredType2, inferredType3, declaredAndInitialized, a, b, c, d, e, f)

	// pointers versus values
	// var valueVariable string = message
	// var pointerVariable *string = &message
	// fmt.Println(message, valueVariable, *pointerVariable)
	// valueVariable += "1"
	// *pointerVariable += "2"
	// fmt.Println(message, valueVariable, *pointerVariable)

	// user defined types cont.
	// var aGreeting greeting.Salutation = "good day"
	// var greetingStruct = greeting.SalutationStruct{"Bob", "hello there"}
	// fmt.Println(aGreeting)
	// fmt.Println(greetingStruct.Greeting, greetingStruct.Name1)

	// constants cont.
	// fmt.Println(greeting.Pi, greeting.Language)
	// fmt.Println(greeting.G, greeting.H, greeting.I)

	// functions cont.
	// greeting.Greet(greetingStruct)
	// greeting.Greet2(greetingStruct)
	// greeting.Greet3(greetingStruct, greeting.Print, false, 1)
	// greeting.Greet3(greetingStruct, greeting.PrintLine, false, 5)
	//greeting.Greet3(salutations, greeting.CreatePrintFunction("!!!"), true, 3)

	salutations.Greet3(greeting.CreatePrintFunction("!!!"), true, 3)
}
