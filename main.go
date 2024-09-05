package main

import (
	"fmt"
	"log"
	"go-learning/getting-started/hello"
	"go-learning/getting-started/greetings"
)

func main() {
	log.SetPrefix("go-learning: ")
	log.SetFlags(0)

	hello.Hello()
	message, err := submod.Message("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
