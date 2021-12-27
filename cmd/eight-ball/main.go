package main

import (
	"fmt"
	"log"
	"net/http"
)

//restful "github.com/emicklei/go-restful/v3"

// TODO: map these to Properties
const (
	Port          = ":8888"
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

	// http.HandleFunc("/", get_default)
	// http.HandleFunc("/random-answer", random_answer)

	// ws := new(restful.WebService)
	// ws.Route(ws.GET("/hello").To(hello))
	// restful.Add(ws)

	log.Fatal(http.ListenAndServe(Port, nil))

	// fmt.Println("Starting server on port " + Port)
	// http.ListenAndServe(":"+Port, nil)
}

// func hello(req *restful.Request, resp *restful.Response) {
// 	io.WriteString(resp, "world")
// }
