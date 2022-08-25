package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	println("SERVER RUNNING!!!")
	http.ListenAndServe(":8000", nil)
}
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Random number is : %v", RandomNumber())
}

// comment here
func RandomNumber() int {
	return rand.Intn(1000)
}
