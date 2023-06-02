package main

import "log"

func main() {
	cat := ""

	if cat == "cat" {
		log.Println("cat is set to cat")
	} else {
		log.Println("cat is not set to cat")
	}

	// switch
	myVar := "bird"

	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")

	case "bird":
		log.Println("myVar is set to bird")

	case "fish":
		log.Println("myVar is set to fish")

	default:
		log.Println("myVar is set to something else")
	}
}
