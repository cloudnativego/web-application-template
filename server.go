package main

import (
	"net/http"
	"text/template"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router) {
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("assets/images"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("assets/css"))))
	mx.HandleFunc("/", homeHandler)
}

type sampleContent struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	data := sampleContent{ID: "8675309", Content: "Hello from Go!"}
	t := template.Must(template.ParseFiles("assets/index.html"))
	t.Execute(w, data)
}
