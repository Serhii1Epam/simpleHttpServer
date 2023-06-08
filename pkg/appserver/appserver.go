// Server package
// simple HTTP server for accepting user request
// check users passwords
package appserver

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Serhii1Epam/simpleHttpServer/pkg/appdb"
	"github.com/Serhii1Epam/simpleHttpServer/pkg/userdata"
)

type Appserver struct {
	Db        *appdb.Database
	User      *userdata.UserData
	is_runned bool
}

const (
	JSON = "application/json"
	TEXT = "text/plain"
)

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
		if err := s.userHandler(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logMsgToWriter(w, r)
			return
		}
		if err := s.User.Login(s.Db); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logMsgToWriter(w, r)
			return
		}
		http.Redirect(w, r, "/accessGaranted", http.StatusMovedPermanently)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		if err := s.userHandler(r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			logMsgToWriter(w, r)
			return
		}
		if err := s.User.Create(s.Db); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logMsgToWriter(w, r)
		}
		w.WriteHeader(http.StatusCreated)
	})
	http.HandleFunc("/accessGaranted", handleAccessGaranted)
	http.HandleFunc("/", handleIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleAccessGaranted(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Congrats!!! Access garanted.)\n")
}

func logMsgToWriter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL [%s] ", r.URL.Path)
	fmt.Fprintf(w, "Method [%v] ", r.Method)
	fmt.Fprintf(w, "Content-Type: \"%s\"\n", r.Header.Get("Content-Type"))
}

func parseRequestBody(r *http.Request) *userdata.UserData {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		return nil
	}
	r.Body.Close()
	switch r.Header.Get("Content-Type") {
	case JSON:
		{
			str := userdata.JsonBytes(req)
			return userdata.Parse(&str)
		}
	case TEXT:
		{
			str := userdata.PlainTextBytes(req)
			return userdata.Parse(&str)
		}
	default:
		{
			return nil
		}
	}
}

func (s Appserver) isCorrectMethod(r *http.Request) bool {
	switch r.Method {
	case http.MethodPost:
		{
			return true
		}
	default:
		{
			return false
		}
	}
}

func (s *Appserver) userHandler(r *http.Request) error {
	if !s.isCorrectMethod(r) {
		customMsg := fmt.Sprintf("Server can't handle method [%v]. Continue...\n", r.Method)
		return errors.New(customMsg)
	}
	if s.User = parseRequestBody(r); s.User == nil {
		return errors.New("Cant parse incoming data")
	}

	return nil
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Simple HTTP Server developed for GO switch program.\n")
	logMsgToWriter(w, r)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	logMsgToWriter(w, r)
}
