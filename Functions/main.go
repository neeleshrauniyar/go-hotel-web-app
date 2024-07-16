package main

import (
	"errors"
	"fmt"
	"net/http"
)

// handler function
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Go")
}

func Rooms(w http.ResponseWriter, r *http.Request) {
	result, _ := addValues(2, 3)
	fmt.Fprintln(w, "Welcome to the Hotel")
	fmt.Fprintln(w, fmt.Sprintf("Welcome to Rooms no: %d", result))
	fmt.Fprintln(w, "Enjoy your stay")
}

func addValues(x, y int) (int, error) {
	var sum int
	sum = x + y
	return sum, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	result, err:= divideValues(2,0)
	if err != nil {
		fmt.Fprintln(w, "Cannot divide by zero")
	}
	fmt.Fprintln(w, "Result: ", result)
}

func divideValues(x, y int) (int, error) {
	if y == 0 {
		err:= errors.New("Cannot divide by zero")
		return 0, err
	}
	result:= x/y
	return result, nil
}

func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/rooms", Rooms)
	http.HandleFunc("/divide", Divide)

	http.ListenAndServe(":8080", nil)
}
