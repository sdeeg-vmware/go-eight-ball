package main

import (
	"fmt"
	"net/http"
)

// TODO: map these to Properties
const (
	Port          = "8888"
	AppName       = "go-eight-ball"
	DefaultAnswer = "Answer is default"
)

//TODO: return html
func get_default(writer http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(writer, "<http><body>"+AppName+": <a href>/random-answer</a></body></http>")
}

func random_answer(writer http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(writer, DefaultAnswer)
}

func main() {

	http.HandleFunc("/", get_default)
	http.HandleFunc("/random-answer", random_answer)

	fmt.Println("Starting server on port " + Port)
	http.ListenAndServe(":"+Port, nil)
}
