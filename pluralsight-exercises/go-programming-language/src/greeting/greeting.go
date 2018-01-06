package greeting

// user defined types
import "fmt"

type Salutation string

type SalutationStruct struct {
	Name1    string
	Greeting string
}
type Printer func(string)

type Salutations []Salutation


// constants
const (
	Pi       = 3.14
	Language = "go"
)
const (
	G = iota
	H
	I
)

// functions
func Greet(s SalutationStruct) {
	fmt.Println("using a function: ")
	fmt.Println(s.Greeting)
	fmt.Println(s.Name1)
}
func Greet2(s SalutationStruct) {
	fmt.Println("using a function with a return with multiple values: ")
	_, alternate := CreateGreetingMessage(s.Greeting, s.Name1)
	// fmt.Println(message)
	fmt.Println(alternate)
	fmt.Println("using a function with a return with multiple NAMED values: ")
	main, backup := CreateGreetingMessageMultiples(s.Greeting, s.Name1)
	fmt.Println(main)
	fmt.Println(backup)
	fmt.Println("using a variadic function")
	mainVariadic, backupVariadic := CreateGreetingMessageVariadic(s.Greeting, s.Name1, "jimmy")
	fmt.Println(mainVariadic)
	fmt.Println(backupVariadic)
}
func CreateGreetingMessage(Greeting, name string) (string, string) {
	return Greeting + ", " + name, "HOWDY, " + name
}

// named multiple returns
func CreateGreetingMessageMultiples(Greeting, name string) (main string, alternate string) {
	main = Greeting + ", " + name
	alternate = "HOWDY, " + name
	return
}

// variadic functions
func CreateGreetingMessageVariadic(Greeting string, name ...string) (mainVariadic, backupVariadic string) {
	mainVariadic = Greeting + ", " + name[0]
	backupVariadic = "HOWDY, " + name[1]
	return
}

// function types and conditionals
func (salutations Salutations)Greet3(do Printer, isFormal bool, times int) {

	for _, h := range salutations {
		main, backup := CreateGreetingMessageMultiples(h.Greeting, h.Name1)
		if prefix := GetPrefix(h.Name1); isFormal {
			do(prefix + main)
		} else {
			do(backup)
		}
	}
}
func Print(s string) {
	fmt.Print(s)
}
func PrintLine(s string) {
	fmt.Println(s)
}

// breaks the signature
// func printCustom(s string, custom string) {
// 	fmt.Println(s + custom)
// }

// use a function creator
func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}

// func GetPrefix(name string) (prefix string) {
// 	switch name {
// 	case "Bob":
// 		prefix = "Mr "
// 		fallthrough
// 	case "Joe", "Amy":
// 		prefix = "Dr "
// 	case "Mary":
// 		prefix = "Mrs "
// 	default:
// 		prefix = "Dude "
// 	}
// 	return
// }

// maps
func GetPrefix(name string) (prefix string) {

	prefixMap := map[string]string{
		"Bob":  "Mr ",
		"Joe":  "Dr ",
		"Amy":  "Dr ",
		"Mary": "Mrs ",
	}

	prefixMap["Joe"] = "Jr "

	delete(prefixMap, "Mary")

	if v, e := prefixMap[name]; e {
		return v
	}
	return "Dude "

}
