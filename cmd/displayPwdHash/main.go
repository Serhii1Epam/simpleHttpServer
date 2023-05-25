// Cmd util for printing password hash
// Use: displayPwdHash <password>
package main

import (
	"fmt"
	"os"

	"github.com/Serhii1Epam/simpleHttpServer/pkg/hasher"
)

func main() {
	var arg string

	if len(os.Args) > 1 {
		arg = os.Args[1]
	} else {
		fmt.Printf("Hashing errors: epmpty password entered.\n")
		fmt.Printf("\tUsage: displayPwdHash <password>")
		return
	}

	hash, err := hasher.HashPassword(arg)

	if err != nil {
		fmt.Println("Hashing errors: arg => ", arg)
		panic("err")
	}

	fmt.Printf("displayPwdHash: Password => %s, Hash => %s", arg, hash)
}
