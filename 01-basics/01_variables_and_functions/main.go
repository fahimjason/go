package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	var whatTOSay string
	var i int

	whatTOSay = "Goodbye cruel world"
	fmt.Println(whatTOSay)

	i = 7
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWasSaid := saySomething()

	fmt.Println("The function said", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "Something", "else"
}
