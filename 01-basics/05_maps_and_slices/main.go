package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// map
	myMap := make(map[string]User)

	user := User{
		FirstName: "Some",
		LastName:  "One",
	}

	myMap["user"] = user

	log.Println(myMap["user"].FirstName)

	// slice
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[6:9])

	srt := []string{"A", "B", "C"}

	log.Println(srt)
}
