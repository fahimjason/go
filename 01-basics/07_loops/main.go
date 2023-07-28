package main

import "log"

func main() {
	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	// loop on slice
	animals := []string{"horse", "cow", "cat", "goat", "ass"}

	for _, animal := range animals {
		log.Println(animal)
	}

	// loop on string
	// var firstLine = "This is the first line"

	// for i, l := range firstLine {
	// 	log.Println(i, ":", l)
	// }

	// loop on map
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
