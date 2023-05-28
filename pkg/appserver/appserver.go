// Server package
// sometext
package appserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Serhii1Epam/simpleHttpServer/pkg/userdata"
)

type Appserver struct {
	db        bool
	is_runned bool
}

func New() *Appserver {
	return &Appserver{db: false, is_runned: false}
}

func (s *Appserver) Run() {
	var err error
	s.db = true
	s.is_runned = true
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/user/login", handleUserLogin)
	http.HandleFunc("/user", handleUser)
	http.HandleFunc("/", handleIndex)

	if err == nil {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL [%s]\n", r.URL.Path)
	fmt.Fprintf(w, "Method [%v]\n", r.Method)
	fmt.Fprintf(w, "Simple HTTP Server developed for GO switch program.\n")
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Process handleUser...\n")

}

func handleUserLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		{
			fmt.Fprintf(w, "Nothing to return in [%v] method. Continue...\n", r.Method)
		}
	case "POST":
		{
			fmt.Fprintf(w, "Process [%v] method...\n", r.Method)
			var jsonVal userdata.UserData
			body, err := ioutil.ReadAll(r.Body)
			if err == nil {
				json.Unmarshal(body, &jsonVal)
			}
			r.Body.Close()
			fmt.Fprintf(w, "User : %s Password : %s\n", jsonVal.User, jsonVal.Password)
		}
	default:
		{
			fmt.Fprintf(w, "Server can't handle method [%v]. Continue...\n", r.Method)
		}
	}

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Continue ...\n")
}
