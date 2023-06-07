// Server package
// simple HTTP server for accepting user request
// check users passwords
package appserver

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Serhii1Epam/simpleHttpServer/pkg/appdb"
	"github.com/Serhii1Epam/simpleHttpServer/pkg/hasher"
	"github.com/Serhii1Epam/simpleHttpServer/pkg/userdata"
)

type Appserver struct {
	Db        *appdb.Database
	User      *userdata.UserData
	Hasher    *hasher.HashingData
	is_runned bool
}

func NewServer() *Appserver {
	return &Appserver{is_runned: false}
}

func (s Appserver) IsRun() bool {
	return s.is_runned
}

func (s *Appserver) SrvRun() {
	s.is_runned = true
	s.Db = appdb.NewDatabase()
	http.HandleFunc("/about", handleAbout)
	http.HandleFunc("/user/login", func(w http.ResponseWriter, r *http.Request) {
		s.userHandler(w, r)
		//fmt.Fprintf(w, "handled User: [%v]\n", s.User)
		s.User.LoginUser(w, s.Db)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		s.userHandler(w, r)
		s.User.CreateUser(w, s.Db)
	})
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

func (s Appserver) checkMethod(r *http.Request) error {
	switch r.Method {
	case http.MethodPost:
		{
			return nil
		}
	default:
		{
			return nil
		}
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	logMsgToWriter(w, r)
	fmt.Fprintf(w, "Simple HTTP Server developed for GO switch program.\n")
}

func (s *Appserver) userHandler(w http.ResponseWriter, r *http.Request) {
	logMsgToWriter(w, r)
	if s.checkMethod(r) != nil {
		//fmt.Fprintf(w, "Server can't handle method [%v]. Continue...\n", r.Method)
		return
	}
	if s.User = parseRequestBody(r); s.User != nil {
		fmt.Fprintf(w, "parseRequestBody [%v]\n", s.User)
	}
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	logMsgToWriter(w, r)
	fmt.Fprintf(w, "Try to use another endpoint.Continue ...\n")
}
