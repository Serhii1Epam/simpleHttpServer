/*
 * Start Simple HTTP server
 */
package main

import (
	"fmt"
)

func main() {
	i := 0
	fmt.Println("Simple HTTP Server start...")
	for {
		i++
		if i == 1000 {
			fmt.Printf("i = %v\n", i)
			break
		}
	}
	fmt.Println("Simple HTTP Server finish work.")
}
