package main

import (
	"fmt"
	"net/http"
)

type Developer struct {
	Name string
	Age  int
}

func FilterUnique(developers []Developer) []string {
	var unique []string
	check := make(map[string]int)
	for _, developer := range developers {
		check[developer.Name] = 1
	}

	for name := range check {
		unique = append(unique, name)
	}
	return unique
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: Homepage")
}

func setup() {
	http.HandleFunc("/", HomePage)
}

var (
	counter = 0
)

func updateCounter() {
	counter++
}
func main() {
	go updateCounter()
	fmt.Println(counter)
	//setup()
	fmt.Println("main method called")
	//log.Fatal(http.ListenAndServe(":10000", nil))

}
