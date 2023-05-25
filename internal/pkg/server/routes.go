// Routes package for simple HTTP Server
package main

import (
	_ "fmt"

	_ "github.com/Serhii1Epam/simpleHttpServer/internal/pkg/server"
)

func (s *server) routes() {
	s.router("/exit", s.handleExit())
	s.router("/about", s.handleAbout())
	s.router("/", s.handleIndex())
}
