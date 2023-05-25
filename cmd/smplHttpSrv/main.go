/*
 * Cmd start Simple HTTP server
 */
package main

import (
	"fmt"
	"log"

	"github.com/Serhii1Epam/simpleHttpServer/internal/pkg/appserver"
)

func main() {
	fmt.Println("Simple HTTP Server start...")
	defer fmt.Println("Simple HTTP Server finish work.")

	srv1 := appserver.Appserver{Is_runned: false}

	log.Fatal(srv1.Run())
}
