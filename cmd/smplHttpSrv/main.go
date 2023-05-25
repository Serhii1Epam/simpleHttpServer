/*
 * Cmd start Simple HTTP server
 */
package main

import (
	"fmt"
	"log"

	_ "github.com/Serhii1Epam/simpleHttpServer/internal/pkg/appServer"
)

func main() {
	fmt.Println("Simple HTTP Server start...")
	defer fmt.Println("Simple HTTP Server finish work.")
	srv := appServer{false}
	log.Fatal(srv.Run())
}
