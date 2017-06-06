package main

import (
	"os"
	"fmt"
	. "github.com/paulohsl/addressline"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Address is missing usage: addressline <address>")
		os.Exit(1)
	}

	res := Parse(os.Args[1])

	if len(res) > 0 {
		fmt.Println(res)
	} else {
		fmt.Println("Cannot parse the Address !")
	}
}
