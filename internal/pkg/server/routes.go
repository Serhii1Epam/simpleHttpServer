package routes

import (
	_ "fmt"
	//_ "github.com/Serhii1Epam/simpleHttpServer/internal/pkg/server"
)

func (s *server) routes() {
	s.router.HandleFunc("/exit", s.handleExit())
	s.router.Get("/about", s.handleAbout())
	s.router.HandleFunc("/", s.handleIndex())
}
