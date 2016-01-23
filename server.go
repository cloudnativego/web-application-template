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
	mx.PathPrefix("/images/").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./assets/images/"))))
	mx.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./assets/css/"))))
	mx.HandleFunc("/", homeHandler)
}

type sampleContent struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

var t *template.Template

func init() {
	t = template.Must(template.ParseFiles("assets/templates/index.html"))
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	data := sampleContent{ID: "8675309", Content: "Hello from Go!"}
	t.Execute(w, data)
}
