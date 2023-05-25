/*
 * Cmd start Simple HTTP server
 */
package main

import (
	"fmt"
	"log"

	_ "github.com/Serhii1Epam/simpleHttpServer/internal/pkg/server"
)

func main() {
	fmt.Println("Simple HTTP Server start...")
	defer fmt.Println("Simple HTTP Server finish work.")
	s := NewServer()

	//http.HandleFunc("/about", aboutHandler)
	//http.HandleFunc("/", indexHandler)
	log.Fatal(s.Run())
}
