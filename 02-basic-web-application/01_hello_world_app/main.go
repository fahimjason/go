package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)

	_, _ = fmt.Fprintf(w, "This is the about page and 2 + 2 is: %d", sum)

}

// addValues add two integer and return the sum
func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, "%s", err)
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 0.0, f)
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cant not divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

// main is the main application function
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	fmt.Println("Number of bytes written:", n)
	// })

	fmt.Printf("Starting application on port %s", portNumber)

	_ = http.ListenAndServe(portNumber, nil)
}
