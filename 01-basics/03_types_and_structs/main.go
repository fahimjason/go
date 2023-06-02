package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName:   "Some",
		LastName:    "One",
		PhoneNumber: "+8801711111111",
	}

	log.Println(user.FirstName, user.LastName, "Birthday:", user.BirthDate)
}
