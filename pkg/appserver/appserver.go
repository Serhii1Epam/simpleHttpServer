// Server package
// sometext
package appserver

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Serhii1Epam/simpleHttpServer/pkg/userdata"
)

type Appserver struct {
	db        bool
	is_runned bool
}

func SrvNew() *Appserver {
	return &Appserver{db: false, is_runned: false}
}

func (s *Appserver) SrvClose() {
	s.is_runned = false
}

func (s *Appserver) SrvRun() {
	s.is_runned = true
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/user/login", handleUserLogin)
	http.HandleFunc("/user", handleUser)
	http.HandleFunc("/", handleIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func logMsgToWriter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL [%s]\n", r.URL.Path)
	fmt.Fprintf(w, "Method [%v]\n", r.Method)
	fmt.Fprintf(w, "Content-Type: \"%s\"\n", r.Header.Get("Content-Type"))
}

func parseRequestBody(r *http.Request) *userdata.UserData {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	r.Body.Close()
	switch r.Header.Get("Content-Type") {
	case "application/json":
		{
			str := userdata.JsonBytes(req)
			return userdata.Parse(&str)
		}
	default:
		{
			str := userdata.UndefinedBytes(req)
			return userdata.Parse(&str)
		}
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	logMsgToWriter(w, r)
	fmt.Fprintf(w, "Simple HTTP Server developed for GO switch program.\n")
}

func handleUser(w http.ResponseWriter, r *http.Request) {
	logMsgToWriter(w, r)
	switch r.Method {
	case "POST":
		{
			reqUserData := parseRequestBody(r)
			reqUserData.LoginUser(w)
		}
	default:
		{
			fmt.Fprintf(w, "Server can't handle method [%v]. Continue...\n", r.Method)
		}
	}
}

func handleUserLogin(w http.ResponseWriter, r *http.Request) {
	logMsgToWriter(w, r)
	switch r.Method {
	case "POST":
		{
			reqUserData := parseRequestBody(r)
			reqUserData.LoginUser(w)
		}
	default:
		{
			fmt.Fprintf(w, "Server can't handle method [%v]. Continue...\n", r.Method)
		}
	}

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	logMsgToWriter(w, r)
	fmt.Fprintf(w, "Try to use another endpoint.Continue ...\n")
}
