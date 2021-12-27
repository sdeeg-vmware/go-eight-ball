package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

// TODO: map these to Properties
const (
	Port    = ":8888"
	AppName = "go-eight-ball"
)

var DefaultAnswer = []string{"Default 0", "Default 1", "Default 2", "Default 3"}

func get_default(req *restful.Request, resp *restful.Response) {

	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(resp, "<http><head></head><body>%s: <a href=\"/random-answer\">random-answer</a></body></http>", AppName)
}

func hello(req *restful.Request, resp *restful.Response) {
	resp.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(resp, "<http><head></head><body>Hello</body></http>")
}

//restful version
func random_answer(req *restful.Request, resp *restful.Response) {

	resp.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(resp, "{ \"answer\": \"%s\" }", DefaultAnswer[rand.Intn(4)])
}

//http version
// func random_answer(writer http.ResponseWriter, req *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	fmt.Fprintf(writer, "{ \"answer\": \"%s\" }", DefaultAnswer)
// }

func main() {

	// http.HandleFunc("/", get_default)
	// http.HandleFunc("/random-answer", random_answer)

	//Using restful package
	ws := new(restful.WebService)
	ws.Route(ws.GET("/").To(get_default))
	ws.Route(ws.GET("/hello").To(hello))
	ws.Route(ws.GET("/random-answer").To(random_answer))
	restful.Add(ws)

	log.Println("Starting server on port " + Port)
	s := &http.Server{
		Addr: Port,
		// Handler:        myHandler,
		// ReadTimeout:    10 * time.Second,
		// WriteTimeout:   10 * time.Second,
		// MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}
