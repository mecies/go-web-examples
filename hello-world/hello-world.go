package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("slucham")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you've requested %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
