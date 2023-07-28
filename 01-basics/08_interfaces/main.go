package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Horse struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	horse := Horse{
		Name:  "Horse",
		Breed: "Arabian",
	}

	printInfo(&horse)

	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "Grey",
		NumberOfTeeth: 38,
	}

	printInfo(&gorilla)
}

func printInfo(a Animal) {
	fmt.Println("The animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}

func (h *Horse) Says() string {
	return "Something about horse"
}

func (h *Horse) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Something about Gorilla"
}

func (h *Gorilla) NumberOfLegs() int {
	return 2
}
