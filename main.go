package main

import (
	"fmt"
	"log"
	"go-learning/getting-started/hello"
	"go-learning/getting-started/greetings"
)

func main() {

	// 1.Hello World
	hello.Hello();
	fmt.Printf("\n");
	
	// 2.Greeting with Parameter
	log.SetPrefix("go-learning: ");
	log.SetFlags(0);
	var message, err = greetings.Message("Sugondese");
	
	if err != nil {
		log.Fatal(err);
	}
	
	fmt.Println(message);
	fmt.Printf("\n");
	// 3. Random greeting
	fmt.Sprintf(greetings.RandomFormat(), "Ligma");

	var greetMap, err = greetings.MultiMessage([]string {"Anderson", "Ligma", "iLadies"});

	if err != nil {
		log.Fatal(err);
	}

	for name, message := range greetMap {
		fmt.Printf(message, name);
		fmt.Printf("\n");
	}
}
