// Server package
// sometext
package appserver

import (
	"fmt"
	"net/http"
	"os"
)

type Appserver struct {
	db        bool
	is_runned bool
}

func (s *Appserver) Run() {
	s.db = true
	http.HandleFunc("/exit", handleExit)
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/", handleIndex)
	s.is_runned = true
	LogFatal(http.ListenAndServe(":8080", nil))
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
