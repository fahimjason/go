package main

import (
	"log"

	"github.com/fahimjason/myprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeTypes
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}
