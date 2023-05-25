/*
 * Cmd start Simple HTTP server
 */
package main

import (
	"fmt"
	"log"

	_ "github.com/github.com/Serhii1Epam/simpleHttpServer/internal/pkg/appserver"
)

func main() {
	fmt.Println("Simple HTTP Server start...")
	defer fmt.Println("Simple HTTP Server finish work.")

	srv := Appserver{is_runned: false}

	log.Fatal(srv.Run())
}
