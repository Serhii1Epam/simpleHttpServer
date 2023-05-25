package server

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/Serhii1Epam/simpleHttpServer/internal/pkg/hasher"
)

type server struct {
	db     *memdb
	router *simpleHttpRouter
}

func newServer() *server {
	s := &server{}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w *http.ResponseWriter, r http.Request) {
	s.router(ServeHTTP(w, r))
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL [%s]\n", r.URL.Path)
	fmt.Fprintf(w, "Method [%v]\n", r.Method)
	fmt.Fprintf(w, "Simple HTTP Server developed for GO switch program.\n")
}

func handleExit(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL [%s]\n", r.URL.Path)
	fmt.Fprintf(w, "Method [%v]\n", r.Method)
	fmt.Fprintf(w, "Closing Simple HTTP Server...\n")
	os.Exit(1)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Process [%v] method...\n", r.Method)
	switch r.Method {
	case "GET":
		{

		}
	case "POST":
		{
			fmt.Fprintf(w, "We shouldn't be here [%v]. Continue...\n", r.Method)
		}
	default:
		{
			fmt.Fprintf(w, "Server can't handle method [%v]. Continue...\n", r.Method)
		}
	}
}
