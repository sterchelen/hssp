package main

import (
	"log"

	"github.com/sterchelen/hssp/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
