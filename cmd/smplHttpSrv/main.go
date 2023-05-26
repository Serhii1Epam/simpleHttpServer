/*
 * Cmd start Simple HTTP server
 */
package main

import (
	"fmt"

	"github.com/Serhii1Epam/simpleHttpServer/pkg/appserver"
	//"simpleHttpServer/pkg/appserver"
)

func main() {
	fmt.Println("Simple HTTP Server start...")
	defer fmt.Println("Simple HTTP Server finish work.")

	srv1 := appserver.Appserver{is_runned: false, db: false}
	srv1.Run()
}
