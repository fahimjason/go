package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello World")

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Number of bytes written:", n)
	})

	_ = http.ListenAndServe(":8080", nil)
}
